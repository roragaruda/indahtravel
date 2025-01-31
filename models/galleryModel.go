package models

import "gorm.io/gorm"

type Gallery struct {
	gorm.Model
	GalleryName string         `gorm:"not null"`
	Photo       []PhotoGallery `gorm:"foreignKey:GalleryID"`
}
