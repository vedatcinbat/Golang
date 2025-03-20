package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Title    string `json:"title"`
	Content  string `json:"content"`
	Price    uint   `json:"price"`
	Quantity uint   `json:"quantity"`
}
