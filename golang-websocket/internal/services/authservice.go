package services

import (
	"encoding/json"
	"fmt"
	"golang-websocket/internal/db"
	"golang-websocket/internal/models"
	"golang-websocket/internal/view"

	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtKey = []byte("supersecret123!@#") // example fixed secret key

type AuthService interface {
	Signup(r *http.Request, w http.ResponseWriter) (view.ApiResponse, error)
	Login(w http.ResponseWriter, r *http.Request)
	Profile(w http.ResponseWriter, r *http.Request) (view.UserResponse, error)
}

type authService struct{}

// Constructor to return interface
func NewAuthService() AuthService {
	return &authService{}
}

func (s *authService) Signup(r *http.Request, w http.ResponseWriter) (view.ApiResponse, error) {

	var apiResponse view.ApiResponse
	var creds models.User
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		apiResponse.Message = "Invalid request"
		return apiResponse, err
	}

	var existingUser models.User
	if err := db.DB.Where("email = ?", creds.Email).First(&existingUser).Error; err == nil {
		// Found a user with this email
		http.Error(w, "User already exists", http.StatusBadRequest)
		apiResponse.Message = "User already existss"
		return apiResponse, err
	} else if err != gorm.ErrRecordNotFound {
		// Some other DB error
		http.Error(w, "Database error", http.StatusInternalServerError)
		apiResponse.Message = "Database error"
		return apiResponse, err
	}
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	creds.Password = string(hashedPassword)

	if err := db.DB.Create(&creds).Error; err != nil {
		http.Error(w, "Failed to save user", http.StatusBadRequest)
		apiResponse.Message = "Failed to save user"
		return apiResponse, err
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Signup successful"})
	apiResponse = view.ApiResponse{
		Message: "Signup successful",
		Status:  http.StatusCreated,
		Data:    nil}

	return apiResponse, nil
}

// Login Handler
func (s *authService) Login(w http.ResponseWriter, r *http.Request) {
	var creds models.User
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var user models.User
	if err := db.DB.Where("email = ?", creds.Email).First(&user).Error; err != nil {
		http.Error(w, "User not found", http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	fmt.Println("Generating token for:", creds.Email)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": creds.Email,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenString, _ := token.SignedString(jwtKey)
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})

}

func (s *authService) Profile(w http.ResponseWriter, r *http.Request) (view.UserResponse, error) {
	// Retrieve email set by middleware
	var userResponse view.UserResponse
	email, ok := r.Context().Value("email").(string)
	if !ok || email == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return userResponse, fmt.Errorf("unauthorized")
	}

	userResponse = view.UserResponse{
		ID:        1,
		Email:     email,
		FirstName: "John",
		LastName:  "Doe",
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(userResponse); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return userResponse, err
	}

	return userResponse, nil
}
