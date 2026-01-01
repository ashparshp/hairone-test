package api

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/hairone/backend/internal/db"
	"github.com/hairone/backend/internal/models"
)

type BookingHandler struct {
	BookingRepo *db.BookingRepository
	ShopRepo    *db.ShopRepository
}

func (h *BookingHandler) CreateBooking(w http.ResponseWriter, r *http.Request) {
	var input models.Booking
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Attach User ID if logged in
	if userID, ok := r.Context().Value(UserIDKey).(int); ok {
		input.UserID = &userID
	}

	// Calculate financials here (simplified port)
	input.FinalPrice = input.TotalPrice // Assuming no discount logic for now
	input.AdminCommission = input.TotalPrice * 0.10
	input.AdminNetRevenue = input.AdminCommission
	input.BarberNetRevenue = input.TotalPrice - input.AdminCommission

	// Generate Random PIN
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	input.BookingKey = fmt.Sprintf("%04d", rng.Intn(10000))

	if err := h.BookingRepo.CreateBooking(r.Context(), &input); err != nil {
		http.Error(w, "Failed to create booking: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(input)
}

func (h *BookingHandler) GetBookings(w http.ResponseWriter, r *http.Request) {
	shopIDStr := r.URL.Query().Get("shopId")
	shopID, err := strconv.Atoi(shopIDStr)
	if err != nil {
		http.Error(w, "Invalid Shop ID", http.StatusBadRequest)
		return
	}

	// Verify Ownership
	userID, ok := r.Context().Value(UserIDKey).(int)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	shop, err := h.ShopRepo.GetShopByID(r.Context(), shopID)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if shop == nil || shop.OwnerID != userID {
		http.Error(w, "You do not own this shop", http.StatusForbidden)
		return
	}

	date := r.URL.Query().Get("date")

	bookings, err := h.BookingRepo.GetBookingsByShop(r.Context(), shopID, date)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(bookings)
}

func (h *BookingHandler) CompleteBooking(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var input struct {
		BookingKey string `json:"bookingKey"`
	}
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Fetch Booking to verify PIN
	booking, err := h.BookingRepo.GetBookingByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if booking == nil {
		http.Error(w, "Booking not found", http.StatusNotFound)
		return
	}

	// Verify PIN
	if booking.BookingKey != input.BookingKey {
		http.Error(w, "Invalid Booking PIN", http.StatusUnauthorized)
		return
	}

	if err := h.BookingRepo.UpdateBookingStatus(r.Context(), id, "completed"); err != nil {
		http.Error(w, "Failed to update status", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
