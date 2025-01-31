package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/attanabilrabbani/indahwisatabe/config"
	"github.com/attanabilrabbani/indahwisatabe/models"
	"github.com/gin-gonic/gin"
)

func GetGallery(c *gin.Context) {
	var galleryData []models.Gallery

	err := config.DB.Preload("Photo").Find(&galleryData).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "unable to receive data at the moment."})
		return
	}

	c.JSON(http.StatusOK, galleryData)

}

func GetGalleryById(c *gin.Context) {
	galleryId := c.Param("id")
	var galleryData models.Gallery

	if err := config.DB.Preload("Photo").First(&galleryData, galleryId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Gallery Not Found.",
		})
		return
	}

	c.JSON(http.StatusOK, galleryData)
}

func EditGalleryName(c *gin.Context) {
	galleryId := c.Param("id")
	var galleryData models.Gallery
	var gallery struct {
		GalleryName string `form:"gallery-name"`
	}

	c.ShouldBind(&gallery)

	err := config.DB.First(&galleryData, galleryId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Update failed. data not found.",
		})
		return
	}

	updateData := make(map[string]interface{})

	if gallery.GalleryName != "" {
		updateData["GalleryName"] = gallery.GalleryName
	}

	if err = config.DB.Model(&galleryData).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update product",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Update Succesful."})
}

func CreateGallery(c *gin.Context) {
	var gallery struct {
		GalleryName string `form:"gallery-name"`
	}

	err := c.ShouldBind(&gallery)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	if gallery.GalleryName == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Name Required"})
		return
	}

	galleryInput := models.Gallery{
		GalleryName: gallery.GalleryName,
	}

	err = config.DB.Create(&galleryInput).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "gallery creation failed"})
	}

	c.JSON(http.StatusOK, gin.H{"ID": galleryInput.ID})
}

func DeleteGallery(c *gin.Context) {
	var gallery models.Gallery
	var galleryPhotos models.PhotoGallery
	galleryId := c.Param("id")

	photosPath := fmt.Sprintf("./assets/image/gallery/%s", galleryId)
	err := os.RemoveAll(photosPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Thumbnail deletion failed",
		})
		return
	}

	err = config.DB.Where("gallery_id=?", galleryId).Delete(&galleryPhotos).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "deletion failed",
		})
	}

	err = config.DB.Where("ID=?", galleryId).Delete(&gallery).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Deletion Failed.",
		})
		return
	}

}
