package routes

import (
	"golang-websocket/internal/handlers"
	"golang-websocket/internal/middleware"
	"net/http"

	"github.com/gorilla/mux"
)
func wrapMiddleware(h http.HandlerFunc, middlewareFunc func(http.Handler) http.Handler) http.HandlerFunc {
	return middlewareFunc(h).ServeHTTP
}
func RegisterRoutes() *mux.Router {
    r := mux.NewRouter()

   // Public routes
	r.HandleFunc("/", handlers.HelloWorldHandler).Methods("GET")
	r.HandleFunc("/ws", handlers.WsHandler).Methods("GET")
	r.HandleFunc("/signup", handlers.Signup).Methods("POST")
	r.HandleFunc("/login", handlers.Login).Methods("POST")

	// Protected routes
	protected := r.PathPrefix("/api").Subrouter()
	protected.Use(middleware.JWTAuth)
	protected.HandleFunc("/profile", handlers.Profile).Methods("GET")


    return r
}
