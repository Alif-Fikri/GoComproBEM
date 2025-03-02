package models

import (
	"time"
)

type User struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `json:"name" validate:"required,min=3"`
	Email     string    `json:"email" gorm:"unique" validate:"required,email"`
	NIM       string    `json:"nim" gorm:"unique" validate:"required,len=11,nim"`
	CreatedAt time.Time `json:"created_at"`
    Reports   []Report  `gorm:"foreignKey:UserID" json:"reports,omitempty"`
}