package handlers

import (
	"encoding/json"
	"fmt"
	"golang-websocket/internal/db"
	"golang-websocket/internal/models"

	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

//var jwtKey = utils.GenerateSecret() // use env var in production
var jwtKey = []byte("supersecret123!@#") // example fixed secret key

// Signup Handler
func Signup(w http.ResponseWriter, r *http.Request) {
	var creds models.User
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(creds.Password), bcrypt.DefaultCost)
	creds.Password = string(hashedPassword)

	if err := db.DB.Create(&creds).Error; err != nil {
		http.Error(w, "User already exists or DB error", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Signup successful"})
}

// Login Handler
func Login(w http.ResponseWriter, r *http.Request) {
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

	// Compare passwords
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(creds.Password)); err != nil {
		http.Error(w, "Invalid password", http.StatusUnauthorized)
		return
	}

	// Generate JWT
	fmt.Println("Generating token for:", creds.Email) // Debugging line
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": creds.Email,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	})
	tokenString, _ := token.SignedString(jwtKey)
fmt.Println("Generated Token:", tokenString) // For debugging; remove in production
	json.NewEncoder(w).Encode(map[string]string{"token": tokenString})
}

func Profile(w http.ResponseWriter, r *http.Request) {
	// Retrieve email set by middleware
	email, ok := r.Context().Value("email").(string)
	if !ok || email == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	fmt.Fprintf(w, "Welcome %s! This is your profile.", email)
}

