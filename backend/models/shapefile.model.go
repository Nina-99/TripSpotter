package models

// import "gorm.io/gorm"

type Shapefile struct {
	// gorm.Model
	Id      uint   `gorm:"primaryKey" json:"id"`
	Name    string `json:"name"`
	GeoJSON string `gorm:"type:json" json:"geojson"`
}
