package routes

import (
	"golang-websocket/internal/handlers"
	"golang-websocket/internal/middleware"
	"golang-websocket/internal/services"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	// Initialize service + handler
	authService := services.NewAuthService()
	authHandler := handlers.NewAuthHandler(authService)

	// Public routes
	r.HandleFunc("/", handlers.HelloWorldHandler).Methods("GET")
	r.HandleFunc("/ws", handlers.WsHandler).Methods("GET")
	r.HandleFunc("/signup", authHandler.Signup).Methods("POST")
	r.HandleFunc("/login", authHandler.Login).Methods("POST")

	// Protected routes
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.JWTAuth)
	protected.HandleFunc("/profile", authHandler.Profile).Methods("GET")

	return r
}
