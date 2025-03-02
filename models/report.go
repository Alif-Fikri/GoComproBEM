package models

import "time"

type Report struct {
    ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
    UserID      uint64    `json:"user_id"`
    Category    string    `json:"category"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Attachment  string    `json:"attachment,omitempty"`
    CreatedAt   time.Time `json:"created_at"`
}
