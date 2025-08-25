package db

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Example DSN: username:password@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
	dsn := "root:password@tcp(127.0.0.1:3306)/db?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}


	fmt.Println("Database connected successfully")
}
