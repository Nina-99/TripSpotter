package models

import "time"

type User struct {
	Id        uint   `gorm:"primary_key" json:"id"`
	Username  string `gorm:"unique;not null" json:"username"`
	Email     string `gorm:"unique;not null" json:"email"`
	Password  string `gorm:"not null" json:"-"`
	Role      string `gorm:"default:user" json:"role"` // "user" or "admin"
	CreatedAt time.Time
	UpdatedAt time.Time
}
