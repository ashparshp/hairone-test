package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/hairone/backend/internal/db"
	"github.com/hairone/backend/internal/models"
)

type ShopHandler struct {
	ShopRepo   *db.ShopRepository
	BarberRepo *db.BarberRepository
}

func (h *ShopHandler) CreateShop(w http.ResponseWriter, r *http.Request) {
	var input models.Shop
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value(UserIDKey).(int)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	input.OwnerID = userID

	if err := h.ShopRepo.CreateShop(r.Context(), &input); err != nil {
		http.Error(w, "Failed to create shop: "+err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(input)
}

func (h *ShopHandler) GetShop(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	shop, err := h.ShopRepo.GetShopByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if shop == nil {
		http.Error(w, "Shop not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(shop)
}

func (h *ShopHandler) CreateBarber(w http.ResponseWriter, r *http.Request) {
	var input models.Barber
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Verify Ownership
	userID, ok := r.Context().Value(UserIDKey).(int)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	shop, err := h.ShopRepo.GetShopByID(r.Context(), input.ShopID)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}
	if shop == nil || shop.OwnerID != userID {
		http.Error(w, "You do not own this shop", http.StatusForbidden)
		return
	}

	if err := h.BarberRepo.CreateBarber(r.Context(), &input); err != nil {
		http.Error(w, "Failed to create barber", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(input)
}

func (h *ShopHandler) GetBarbers(w http.ResponseWriter, r *http.Request) {
	shopIDStr := r.URL.Query().Get("shopId")
	shopID, err := strconv.Atoi(shopIDStr)
	if err != nil {
		http.Error(w, "Invalid Shop ID", http.StatusBadRequest)
		return
	}

	barbers, err := h.BarberRepo.GetBarbersByShop(r.Context(), shopID)
	if err != nil {
		http.Error(w, "Database error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(barbers)
}
