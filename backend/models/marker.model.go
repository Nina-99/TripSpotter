package models

import "gorm.io/gorm"

type Marker struct {
	gorm.Model
	Id     uint   `gorm:"primaryKey"`
	Name   string `json:"name"`
	ShpId  uint
	Images []Image
}
