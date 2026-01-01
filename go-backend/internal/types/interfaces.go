package types

import (
	"context"
	"github.com/hairone/backend/internal/models"
)

type UserRepository interface {
	CreateUser(ctx context.Context, u *models.User) error
	GetUserByPhone(ctx context.Context, phone string) (*models.User, error)
	GetUserByID(ctx context.Context, id int) (*models.User, error)
}

type ShopRepository interface {
	CreateShop(ctx context.Context, s *models.Shop) error
	GetShopByID(ctx context.Context, id int) (*models.Shop, error)
}

type BarberRepository interface {
	CreateBarber(ctx context.Context, b *models.Barber) error
	GetBarbersByShop(ctx context.Context, shopID int) ([]models.Barber, error)
}

type BookingRepository interface {
	CreateBooking(ctx context.Context, b *models.Booking) error
	GetBookingsByShop(ctx context.Context, shopID int, date string) ([]models.Booking, error)
	GetBookingByID(ctx context.Context, id int) (*models.Booking, error)
	UpdateBookingStatus(ctx context.Context, id int, status string) error
}
