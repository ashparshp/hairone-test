package models

import (
	"time"
)

type User struct {
	ID                int       `json:"id"`
	Phone             string    `json:"phone"`
	PasswordHash      string    `json:"-"`
	Password          string    `json:"password,omitempty"` // Input only
	Name              string    `json:"name"`
	Avatar            string    `json:"avatar"`
	Email             string    `json:"email"`
	Gender            string    `json:"gender"`
	Role              string    `json:"role"` // 'user', 'admin', 'owner'
	IsPremium         bool      `json:"isPremium"`
	BusinessName      string    `json:"businessName"`
	ApplicationStatus string    `json:"applicationStatus"`
	SuspensionReason  string    `json:"suspensionReason"`
	TokenVersion      int       `json:"tokenVersion"`
	CreatedAt         time.Time `json:"createdAt"`
}

type Shop struct {
	ID                  int         `json:"id"`
	OwnerID             int         `json:"ownerId"`
	Name                string      `json:"name"`
	Address             string      `json:"address"`
	Coordinates         interface{} `json:"coordinates"` // JSONB
	Image               string      `json:"image"`
	Gallery             interface{} `json:"gallery"` // JSONB
	Type                string      `json:"type"`
	Rating              float64     `json:"rating"`
	ReviewCount         int         `json:"reviewCount"`
	BufferTime          int         `json:"bufferTime"`
	MinBookingNotice    int         `json:"minBookingNotice"`
	MaxBookingNotice    int         `json:"maxBookingNotice"`
	AutoApproveBookings bool        `json:"autoApproveBookings"`
	BlockCustomBookings bool        `json:"blockCustomBookings"`
	IsDisabled          bool        `json:"isDisabled"`
	Services            []Service   `json:"services,omitempty"` // Loaded separately
	Combos              []Combo     `json:"combos,omitempty"`   // Loaded separately
}

type Service struct {
	ID          int     `json:"id"`
	ShopID      int     `json:"shopId"`
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Duration    int     `json:"duration"`
	IsAvailable bool    `json:"isAvailable"`
}

type Combo struct {
	ID            int         `json:"id"`
	ShopID        int         `json:"shopId"`
	Name          string      `json:"name"`
	Price         float64     `json:"price"`
	OriginalPrice float64     `json:"originalPrice"`
	Duration      int         `json:"duration"`
	Items         interface{} `json:"items"` // JSONB
	IsAvailable   bool        `json:"isAvailable"`
}

type Barber struct {
	ID             int         `json:"id"`
	ShopID         int         `json:"shopId"`
	Name           string      `json:"name"`
	Avatar         string      `json:"avatar"`
	StartHour      string      `json:"startHour"`
	EndHour        string      `json:"endHour"`
	Breaks         interface{} `json:"breaks"`         // JSONB
	WeeklySchedule interface{} `json:"weeklySchedule"` // JSONB
	SpecialHours   interface{} `json:"specialHours"`   // JSONB
	IsAvailable    bool        `json:"isAvailable"`
}

type Booking struct {
	ID               int         `json:"id"`
	UserID           *int        `json:"userId"`
	ShopID           int         `json:"shopId"`
	BarberID         int         `json:"barberId"`
	ServiceNames     interface{} `json:"serviceNames"` // JSONB
	TotalPrice       float64     `json:"totalPrice"`
	TotalDuration    int         `json:"totalDuration"`
	Date             string      `json:"date"`
	StartTime        string      `json:"startTime"`
	EndTime          string      `json:"endTime"`
	Status           string      `json:"status"`
	Type             string      `json:"type"`
	PaymentMethod    string      `json:"paymentMethod"`
	BookingKey       string      `json:"bookingKey"`
	IsRated          bool        `json:"isRated"`
	Notes            string      `json:"notes"`
	OriginalPrice    float64     `json:"originalPrice"`
	DiscountAmount   float64     `json:"discountAmount"`
	FinalPrice       float64     `json:"finalPrice"`
	AdminCommission  float64     `json:"adminCommission"`
	AdminNetRevenue  float64     `json:"adminNetRevenue"`
	BarberNetRevenue float64     `json:"barberNetRevenue"`
	AmountCollectedBy string     `json:"amountCollectedBy"`
	SettlementStatus string      `json:"settlementStatus"`
	SettlementID     *int        `json:"settlementId"`
	CreatedAt        time.Time   `json:"createdAt"`
}
