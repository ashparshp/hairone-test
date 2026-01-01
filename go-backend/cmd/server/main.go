package main

import (
	"log"
	"net/http"
	"time"

	"github.com/hairone/backend/internal/api"
	"github.com/hairone/backend/internal/db"
	"github.com/hairone/backend/internal/service"
)

func main() {
	if err := db.Connect(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	if err := db.Migrate("internal/db/schema.sql"); err != nil {
		log.Printf("Migration warning: %v", err)
	}

	userRepo := &db.UserRepository{}
	shopRepo := &db.ShopRepository{}
	barberRepo := &db.BarberRepository{}
	bookingRepo := &db.BookingRepository{}

	authHandler := &api.AuthHandler{UserRepo: userRepo}
	shopHandler := &api.ShopHandler{ShopRepo: shopRepo, BarberRepo: barberRepo}
	bookingHandler := &api.BookingHandler{BookingRepo: bookingRepo, ShopRepo: shopRepo}

	mux := http.NewServeMux()

	// Auth Routes
	mux.HandleFunc("POST /api/auth/register", authHandler.Register)
	mux.HandleFunc("POST /api/auth/login", authHandler.Login)

	// Public Routes
	mux.HandleFunc("GET /api/shops/{id}", shopHandler.GetShop)
	mux.HandleFunc("GET /api/barbers", shopHandler.GetBarbers)

	// Protected Routes
	mux.Handle("POST /api/shops", api.AuthMiddleware(http.HandlerFunc(shopHandler.CreateShop)))
	mux.Handle("POST /api/barbers", api.AuthMiddleware(http.HandlerFunc(shopHandler.CreateBarber)))
	mux.Handle("GET /api/bookings", api.AuthMiddleware(http.HandlerFunc(bookingHandler.GetBookings)))
	mux.Handle("POST /api/bookings", api.AuthMiddleware(http.HandlerFunc(bookingHandler.CreateBooking)))
	mux.Handle("POST /api/bookings/{id}/complete", api.AuthMiddleware(http.HandlerFunc(bookingHandler.CompleteBooking)))

	// Settlement Job (Background Ticker)
	go func() {
		settlementService := &service.SettlementService{}
		// Run once on startup for demo
		// settlementService.RunSettlementJob()

		// Schedule daily (simplified)
		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()
		for range ticker.C {
			settlementService.RunSettlementJob()
		}
	}()

	// Protected Route Example
	mux.Handle("GET /api/me", api.AuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{"message": "You are authorized"}`))
	})))

	// CORS Middleware
	corsHandler := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}

			next.ServeHTTP(w, r)
		})
	}

	log.Println("Server starting on :8080")
	if err := http.ListenAndServe(":8080", corsHandler(mux)); err != nil {
		log.Fatal(err)
	}
}
