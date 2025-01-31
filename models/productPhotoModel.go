package models

import "gorm.io/gorm"

type ProductPhoto struct {
	gorm.Model
	ImageName string
	ProductID uint `gorm:"not null"`
}
