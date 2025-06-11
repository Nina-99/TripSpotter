package models

import "time"

type Review struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	SiteID    uint      `json:"site_id"`
	Stars     int       `json:"stars"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}
