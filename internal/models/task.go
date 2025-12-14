package models

import "time"

type Task struct {
	ID        uint   `gorm:"primaryKey"`
	Title     string `gorm:"not null"`
	Completed bool   `gorm:"default:false"`
	UserID    uint   `gorm:"index;not null"`
	CreatedAt time.Time
}
