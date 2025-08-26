package migration

import (
	"golang-websocket/internal/db"
	"golang-websocket/internal/models"
	"log"
)

func Migrate() {
	if db.DB == nil {
		log.Fatal("DB connection not initialized")
	}

	if err := db.DB.AutoMigrate(&models.User{}, &models.Product{}); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	log.Println("Database migration completed")
}
