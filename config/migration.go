package config

import "github.com/attanabilrabbani/indahwisatabe/models"

func MigrateDB() {
	DB.AutoMigrate(&models.Product{}, &models.Admin{}, &models.Order{}, &models.Gallery{}, &models.PhotoGallery{},
		&models.ProductPhoto{})
}
