package models

type Image struct {
	Id       uint `gorm:"primaryKey"`
	Filename string
	Path     string
	MarkerId uint
}
