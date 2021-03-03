package models

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          string  `gorm:"primaryKey autoIncrement" json:"id"`
	Title       string  `gorm:"type:varchar(30) not null" json:"title"`
	Description string  `gorm:"type:varchar(300)" json:"description"`
	Price       float64 `json:"price"`
}
