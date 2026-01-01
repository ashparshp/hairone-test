package db

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hairone/backend/internal/models"
	"github.com/jackc/pgx/v5"
)

type ShopRepository struct{}

func (r *ShopRepository) CreateShop(ctx context.Context, s *models.Shop) error {
	tx, err := Pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	coordsJSON, _ := json.Marshal(s.Coordinates)
	galleryJSON, _ := json.Marshal(s.Gallery)

	query := `
		INSERT INTO shops (owner_id, name, address, coordinates, image, gallery, type)
		VALUES (@ownerId, @name, @address, @coordinates, @image, @gallery, @type)
		RETURNING id
	`
	args := pgx.NamedArgs{
		"ownerId":     s.OwnerID,
		"name":        s.Name,
		"address":     s.Address,
		"coordinates": coordsJSON,
		"image":       s.Image,
		"gallery":     galleryJSON,
		"type":        s.Type,
	}

	err = tx.QueryRow(ctx, query, args).Scan(&s.ID)
	if err != nil {
		return fmt.Errorf("failed to insert shop: %w", err)
	}

	// Insert Services if any
	for _, svc := range s.Services {
		_, err := tx.Exec(ctx, `
			INSERT INTO services (shop_id, name, price, duration)
			VALUES ($1, $2, $3, $4)
		`, s.ID, svc.Name, svc.Price, svc.Duration)
		if err != nil {
			return fmt.Errorf("failed to insert service: %w", err)
		}
	}

	return tx.Commit(ctx)
}

func (r *ShopRepository) GetShopByID(ctx context.Context, id int) (*models.Shop, error) {
	query := `SELECT id, owner_id, name, address, coordinates, image, gallery, type, rating FROM shops WHERE id = $1`
	var s models.Shop
	var coordsBytes, galleryBytes []byte

	err := Pool.QueryRow(ctx, query, id).Scan(&s.ID, &s.OwnerID, &s.Name, &s.Address, &coordsBytes, &s.Image, &galleryBytes, &s.Type, &s.Rating)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if len(coordsBytes) > 0 {
		json.Unmarshal(coordsBytes, &s.Coordinates)
	}
	if len(galleryBytes) > 0 {
		json.Unmarshal(galleryBytes, &s.Gallery)
	}

	// Load Services (Basic example, better to use Join or separate query)
	rows, err := Pool.Query(ctx, `SELECT id, name, price, duration FROM services WHERE shop_id = $1`, s.ID)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var svc models.Service
			if err := rows.Scan(&svc.ID, &svc.Name, &svc.Price, &svc.Duration); err == nil {
				svc.ShopID = s.ID
				s.Services = append(s.Services, svc)
			}
		}
	}

	return &s, nil
}

type BarberRepository struct{}

func (r *BarberRepository) CreateBarber(ctx context.Context, b *models.Barber) error {
	breaksJSON, _ := json.Marshal(b.Breaks)
	scheduleJSON, _ := json.Marshal(b.WeeklySchedule)
	specialJSON, _ := json.Marshal(b.SpecialHours)

	query := `
		INSERT INTO barbers (shop_id, name, avatar, start_hour, end_hour, breaks, weekly_schedule, special_hours)
		VALUES (@shopId, @name, @avatar, @startHour, @endHour, @breaks, @weeklySchedule, @specialHours)
		RETURNING id
	`
	args := pgx.NamedArgs{
		"shopId":         b.ShopID,
		"name":           b.Name,
		"avatar":         b.Avatar,
		"startHour":      b.StartHour,
		"endHour":        b.EndHour,
		"breaks":         breaksJSON,
		"weeklySchedule": scheduleJSON,
		"specialHours":   specialJSON,
	}
	return Pool.QueryRow(ctx, query, args).Scan(&b.ID)
}

func (r *BarberRepository) GetBarbersByShop(ctx context.Context, shopID int) ([]models.Barber, error) {
	query := `SELECT id, shop_id, name, avatar, start_hour, end_hour, weekly_schedule FROM barbers WHERE shop_id = $1`
	rows, err := Pool.Query(ctx, query, shopID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var barbers []models.Barber
	for rows.Next() {
		var b models.Barber
		var scheduleBytes []byte
		if err := rows.Scan(&b.ID, &b.ShopID, &b.Name, &b.Avatar, &b.StartHour, &b.EndHour, &scheduleBytes); err != nil {
			return nil, err
		}
		if len(scheduleBytes) > 0 {
			json.Unmarshal(scheduleBytes, &b.WeeklySchedule)
		}
		barbers = append(barbers, b)
	}
	return barbers, nil
}
