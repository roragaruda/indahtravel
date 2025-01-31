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

func UploadProductPhoto(c *gin.Context) {
	productID, _ := strconv.Atoi(c.Param("id"))

	form, _ := c.MultipartForm()
	files := form.File["photos"]

	if len(files) != 0 {
		filePath := fmt.Sprintf("./assets/image/productphoto/%d", productID)
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

			productPhoto := models.ProductPhoto{
				ImageName: imageName,
				ProductID: uint(productID),
			}

			err = config.DB.Create(&productPhoto).Error
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "error uploading images",
				})
			}

		}
		c.JSON(http.StatusOK, gin.H{"message": "Images uploaded successfully"})
	}
}

func DeletePhoto(c *gin.Context) {
	var photo models.ProductPhoto
	photoId := c.Param("id")

	result := config.DB.Where("ID=?", photoId).First(&photo)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "not found"})
		return
	}

	photoPath := fmt.Sprintf("./assets/image/productphoto/%d/%s", photo.ProductID, photo.ImageName)
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

// func UploadProductPhoto(c *gin.Context) {
// 	productId := c.Param("id")
// 	newProductId, _ := strconv.Atoi(productId)

// 	var input struct {
// 		ImageName string
// 		ProductId uint
// 	}

// 	err := c.ShouldBind(&input)

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "invalid entry",
// 		})
// 		return
// 	}

// 	// form, err := c.MultipartForm()
// 	// if err != nil {
// 	// 	c.JSON(http.StatusBadRequest, gin.H{
// 	// 		"error": "Failed to parse form",
// 	// 	})
// 	// 	return
// 	// }

// 	filePath := fmt.Sprintf("./assets/image/productphoto/%d", newProductId)
// 	if _, err := os.Stat(filePath); os.IsNotExist(err) {
// 		err := os.MkdirAll(filePath, os.ModePerm)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error": "Unable to create directory",
// 			})
// 			return
// 		}
// 	}

// 	productPhoto := models.ProductPhoto{
// 		ImageName: input.ImageName,
// 		ProductID: uint(newProductId),
// 	}

// 	err = config.DB.Create(&productPhoto).Error
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "invalid entry",
// 		})
// 		return
// 	}

// 	photos, err := c.FormFile("photos")
// 	if err == nil {
// 		imageName := strings.ReplaceAll(photos.Filename, " ", "_")
// 		imagePath := fmt.Sprintf("%s/%s", filePath, imageName)
// 		err = c.SaveUploadedFile(photos, imagePath)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error": "Failed to add thumbnail.",
// 			})
// 			return
// 		}
// 		productPhoto.ImageName = imageName
// 		err = config.DB.Save(&productPhoto).Error
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{
// 				"error": "failed to save product.",
// 			})
// 			return
// 		}
// 	}

// 	// for _, file := range files {
// 	// 	imageName := strings.ReplaceAll(file.Filename, " ", "_")
// 	// 	imgPath := fmt.Sprintf("%s/%s", filePath, imageName)
// 	// 	if err := c.SaveUploadedFile(file, imgPath); err != nil {
// 	// 		c.JSON(http.StatusInternalServerError, gin.H{
// 	// 			"error": "error Uploading image",
// 	// 		})
// 	// 		return
// 	// 	}

// 	// 	productPhoto := models.ProductPhoto{
// 	// 		ImageName: imageName,
// 	// 		ProductID: uint(newProductId),
// 	// 	}

// 	// 	fmt.Println(productPhoto.ImageName)

// 	// 	err := config.DB.Create(&productPhoto).Error

// 	// 	if err != nil {
// 	// 		c.JSON(http.StatusInternalServerError, gin.H{
// 	// 			"error": "failed to upload photos",
// 	// 		})
// 	// 		return
// 	// 	}

// 	// }

// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "uploads succesful",
// 	})

// }
