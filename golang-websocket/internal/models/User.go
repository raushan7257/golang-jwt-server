package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email     string `gorm:"unique" json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
