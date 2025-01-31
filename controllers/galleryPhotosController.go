package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/attanabilrabbani/indahwisatabe/config"
	"github.com/attanabilrabbani/indahwisatabe/models"
	"github.com/gin-gonic/gin"
)

func UploadGalleryPhotos(c *gin.Context) {
	galleryID, _ := strconv.Atoi(c.Param("id"))

	form, _ := c.MultipartForm()
	files := form.File["gallery-photos"]

	if len(files) != 0 {
		filePath := fmt.Sprintf("./assets/image/gallery/%d", galleryID)
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			err := os.MkdirAll(filePath, os.ModePerm)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Unable to create directory",
				})
				return
			}
		}
		for _, photos := range files {
			imageName := strings.ReplaceAll(photos.Filename, " ", "_")
			imagePath := fmt.Sprintf("%s/%s", filePath, imageName)

			err := c.SaveUploadedFile(photos, imagePath)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "failed to save file",
				})
			}

			galleryPhoto := models.PhotoGallery{
				FotoGallery: imageName,
				GalleryID:   uint(galleryID),
			}

			err = config.DB.Create(&galleryPhoto).Error
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "error uploading images",
				})
			}

		}
		c.JSON(http.StatusOK, gin.H{"message": "Images uploaded successfully"})
	}
}

func DeleteGalleryPhotos(c *gin.Context) {
	var photo models.PhotoGallery
	photoId := c.Param("id")

	result := config.DB.Where("ID=?", photoId).First(&photo)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "not found"})
		return
	}

	photoPath := fmt.Sprintf("./assets/image/gallery/%d/%s", photo.GalleryID, photo.FotoGallery)
	fmt.Println(photoPath)
	err := os.Remove(photoPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Photo deletion failed",
		})
		return
	}
	err = config.DB.Delete(&photo).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Deletion Failed.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product Deleted Succesfully.",
	})

}
