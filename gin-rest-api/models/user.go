package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"first_name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}
