package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uuid.UUID `gorm:"type:char(36);primaryKey;" json:"id"`
	Username     string    `gorm:"not null" json:"username"`
	PassHash     string    `gorm:"not null" json:"hash"`
	Salt         string    `gorm:"not null" json:"salt"`
	EmailAddress *string   `gorm:"not null unique_index" json:"email"`
}
