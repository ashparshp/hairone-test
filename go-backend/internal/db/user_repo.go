package db

import (
	"context"
	"fmt"
	"github.com/hairone/backend/internal/models"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct{}

func (r *UserRepository) CreateUser(ctx context.Context, u *models.User) error {
	query := `
		INSERT INTO users (phone, password_hash, name, avatar, email, gender, role, is_premium, business_name, application_status, suspension_reason)
		VALUES (@phone, @passwordHash, @name, @avatar, @email, @gender, @role, @isPremium, @businessName, @applicationStatus, @suspensionReason)
		RETURNING id, created_at
	`
	args := pgx.NamedArgs{
		"phone":             u.Phone,
		"passwordHash":      u.PasswordHash,
		"name":              u.Name,
		"avatar":            u.Avatar,
		"email":             u.Email,
		"gender":            u.Gender,
		"role":              u.Role,
		"isPremium":         u.IsPremium,
		"businessName":      u.BusinessName,
		"applicationStatus": u.ApplicationStatus,
		"suspensionReason":  u.SuspensionReason,
	}

	err := Pool.QueryRow(ctx, query, args).Scan(&u.ID, &u.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

func (r *UserRepository) GetUserByPhone(ctx context.Context, phone string) (*models.User, error) {
	query := `SELECT id, phone, password_hash, name, avatar, email, role, is_premium FROM users WHERE phone = $1`
	var u models.User
	err := Pool.QueryRow(ctx, query, phone).Scan(&u.ID, &u.Phone, &u.PasswordHash, &u.Name, &u.Avatar, &u.Email, &u.Role, &u.IsPremium)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id int) (*models.User, error) {
	query := `SELECT id, phone, name, avatar, email, role, is_premium FROM users WHERE id = $1`
	var u models.User
	err := Pool.QueryRow(ctx, query, id).Scan(&u.ID, &u.Phone, &u.Name, &u.Avatar, &u.Email, &u.Role, &u.IsPremium)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	return &u, nil
}
