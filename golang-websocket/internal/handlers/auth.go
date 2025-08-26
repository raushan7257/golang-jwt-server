package handlers

import (
	"golang-websocket/internal/services"

	"net/http"
)

type AuthHandler struct {
	Service services.AuthService
}

func NewAuthHandler(service services.AuthService) *AuthHandler {
	return &AuthHandler{
		Service: service,
	}
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	h.Service.Signup(r, w)
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	h.Service.Login(w, r)
}

func (h *AuthHandler) Profile(w http.ResponseWriter, r *http.Request) {
	h.Service.Profile(w, r)
}
