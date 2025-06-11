package models

type Shapefile struct {
	Id      uint   `gorm:"primaryKey" json:"id"`
	Name    string `json:"name"`
	GeoJSON string `gorm:"type:json" json:"geojson"`
}
