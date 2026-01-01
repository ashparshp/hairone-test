package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/hairone/backend/internal/db"
)

type SettlementService struct{}

func (s *SettlementService) RunSettlementJob() {
	ctx := context.Background()
	log.Println("--- STARTING SETTLEMENT JOB ---")

	// 1. Define Cutoff (Start of current week Monday)
	// In Go, finding "Start of Week" needs manual calculation
	now := time.Now()
	weekday := int(now.Weekday())
	if weekday == 0 {
		weekday = 7 // Make Sunday = 7 to calculate Monday offset easily
	}
	daysToSubtract := weekday - 1
	cutoffDate := now.AddDate(0, 0, -daysToSubtract).Truncate(24 * time.Hour) // Monday 00:00

	cutoffDateStr := cutoffDate.Format("2006-01-02")
	log.Printf("Searching for unsettled completed bookings before: %s", cutoffDateStr)

	// 2. Aggregation Query (SQL version of the Mongo pipeline)
	// Group by Shop, Filter by Completed & Unsettled & Date < Cutoff
	// Sum AdminNetRevenue where Payment = 'CASH' (Shop owes Admin)
	// Sum BarberNetRevenue where Payment != 'CASH' (Admin owes Shop)

	query := `
		SELECT
			shop_id,
			array_agg(id) as booking_ids,
			MIN(date) as min_date,
			MAX(date) as max_date,
			SUM(CASE WHEN payment_method = 'CASH' THEN admin_net_revenue ELSE 0 END) as total_admin_net, -- Shop owes Admin
			SUM(CASE WHEN payment_method != 'CASH' THEN barber_net_revenue ELSE 0 END) as total_barber_net -- Admin owes Shop
		FROM bookings
		WHERE status = 'completed'
		  AND (settlement_status = 'PENDING' OR settlement_status IS NULL)
		  AND date < $1
		GROUP BY shop_id
	`

	rows, err := db.Pool.Query(ctx, query, cutoffDateStr)
	if err != nil {
		log.Printf("Error querying settlements: %v", err)
		return
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		var shopID int
		var bookingIDs []int
		var minDate, maxDate string
		var totalAdminNet, totalBarberNet float64

		// Note: array_agg returns []any in pgx unless typed, scanning into []int works
		if err := rows.Scan(&shopID, &bookingIDs, &minDate, &maxDate, &totalAdminNet, &totalBarberNet); err != nil {
			log.Printf("Error scanning settlement row: %v", err)
			continue
		}

		// Calculate Net
		// Admin Owes Shop = totalBarberNet
		// Shop Owes Admin = totalAdminNet
		// Net (Payout to Shop) = totalBarberNet - totalAdminNet

		netAmount := totalBarberNet - totalAdminNet
		settlementType := "PAYOUT"
		finalAmount := netAmount

		if netAmount < 0 {
			settlementType = "COLLECTION"
			finalAmount = -netAmount
		}

		status := "PENDING_PAYOUT"
		if settlementType == "COLLECTION" {
			status = "PENDING_COLLECTION"
		}

		// Transaction to Create Settlement and Update Bookings
		if err := createSettlement(ctx, shopID, settlementType, finalAmount, status, bookingIDs, minDate, maxDate); err != nil {
			log.Printf("Error creating settlement for shop %d: %v", shopID, err)
		} else {
			count++
		}
	}

	log.Printf("--- SETTLEMENT JOB COMPLETE: Processed %d shops ---", count)
}

func createSettlement(ctx context.Context, shopID int, sType string, amount float64, status string, bookingIDs []int, start, end string) error {
	tx, err := db.Pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// Insert Settlement
	var settlementID int
	bookingsJSON, _ := json.Marshal(bookingIDs)

	err = tx.QueryRow(ctx, `
		INSERT INTO settlements (shop_id, type, amount, status, bookings, date_range_start, date_range_end, notes)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id
	`, shopID, sType, amount, status, bookingsJSON, start, end, "Auto-generated via Go Cron").Scan(&settlementID)

	if err != nil {
		return fmt.Errorf("insert settlement: %w", err)
	}

	// Update Bookings
	_, err = tx.Exec(ctx, `
		UPDATE bookings
		SET settlement_status = 'SETTLED', settlement_id = $1
		WHERE id = ANY($2)
	`, settlementID, bookingIDs)

	if err != nil {
		return fmt.Errorf("update bookings: %w", err)
	}

	return tx.Commit(ctx)
}
