package db

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hairone/backend/internal/models"
	"github.com/jackc/pgx/v5"
)

type BookingRepository struct{}

func (r *BookingRepository) CreateBooking(ctx context.Context, b *models.Booking) error {
	servicesJSON, _ := json.Marshal(b.ServiceNames)

	query := `
		INSERT INTO bookings (
			user_id, shop_id, barber_id, service_names, total_price, total_duration,
			date, start_time, end_time, status, type, payment_method, booking_key,
			original_price, discount_amount, final_price, admin_commission, admin_net_revenue, barber_net_revenue
		)
		VALUES (
			@userId, @shopId, @barberId, @serviceNames, @totalPrice, @totalDuration,
			@date, @startTime, @endTime, @status, @type, @paymentMethod, @bookingKey,
			@originalPrice, @discountAmount, @finalPrice, @adminCommission, @adminNetRevenue, @barberNetRevenue
		)
		RETURNING id, created_at
	`
	args := pgx.NamedArgs{
		"userId":           b.UserID,
		"shopId":           b.ShopID,
		"barberId":         b.BarberID,
		"serviceNames":     servicesJSON,
		"totalPrice":       b.TotalPrice,
		"totalDuration":    b.TotalDuration,
		"date":             b.Date,
		"startTime":        b.StartTime,
		"endTime":          b.EndTime,
		"status":           b.Status,
		"type":             b.Type,
		"paymentMethod":    b.PaymentMethod,
		"bookingKey":       b.BookingKey,
		"originalPrice":    b.OriginalPrice,
		"discountAmount":   b.DiscountAmount,
		"finalPrice":       b.FinalPrice,
		"adminCommission":  b.AdminCommission,
		"adminNetRevenue":  b.AdminNetRevenue,
		"barberNetRevenue": b.BarberNetRevenue,
	}

	err := Pool.QueryRow(ctx, query, args).Scan(&b.ID, &b.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to create booking: %w", err)
	}
	return nil
}

func (r *BookingRepository) GetBookingsByShop(ctx context.Context, shopID int, date string) ([]models.Booking, error) {
	// Simple filter by date if provided
	query := `
		SELECT id, user_id, shop_id, barber_id, date, start_time, end_time, status, total_price
		FROM bookings
		WHERE shop_id = $1 AND ($2 = '' OR date = $2)
		ORDER BY date, start_time
	`
	rows, err := Pool.Query(ctx, query, shopID, date)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var bookings []models.Booking
	for rows.Next() {
		var b models.Booking
		if err := rows.Scan(&b.ID, &b.UserID, &b.ShopID, &b.BarberID, &b.Date, &b.StartTime, &b.EndTime, &b.Status, &b.TotalPrice); err != nil {
			return nil, err
		}
		bookings = append(bookings, b)
	}
	return bookings, nil
}

func (r *BookingRepository) GetBookingByID(ctx context.Context, id int) (*models.Booking, error) {
	query := `SELECT id, user_id, shop_id, booking_key, status FROM bookings WHERE id = $1`
	var b models.Booking
	err := Pool.QueryRow(ctx, query, id).Scan(&b.ID, &b.UserID, &b.ShopID, &b.BookingKey, &b.Status)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &b, nil
}

func (r *BookingRepository) UpdateBookingStatus(ctx context.Context, id int, status string) error {
	_, err := Pool.Exec(ctx, `UPDATE bookings SET status = $1 WHERE id = $2`, status, id)
	return err
}
