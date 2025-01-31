package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	NamaDestinasi string    `gorm:"not null" form:"nama-destinasi"`
	Harga         int       `gorm:"not null" form:"harga"`
	Deskripsi     string    `form:"deskripsi"`
	StartDate     time.Time `binding:"required" gorm:"type:date"`
	EndDate       time.Time `binding:"required" gorm:"type:date"`
	Thumbnail     string
	Order         []Order `gorm:"foreignKey:ProductID"`
	// Date          ProductDate    `gorm:"foreignKey:ProductID"`
	Photo []ProductPhoto `gorm:"foreignKey:ProductID"`
}
