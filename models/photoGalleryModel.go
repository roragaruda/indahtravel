package models

import "gorm.io/gorm"

type PhotoGallery struct {
	gorm.Model
	FotoGallery string `gorm:"not null"`
	GalleryID   uint   `gorm:"not null"`
}
