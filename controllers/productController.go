package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/attanabilrabbani/indahwisatabe/config"
	"github.com/attanabilrabbani/indahwisatabe/models"
	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {
	var products []models.Product

	err := config.DB.Preload("Photo").Preload("Order").Find(&products).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to ",
		})
		return
	}

	c.JSON(http.StatusOK, products)
}

func GetProductById(c *gin.Context) {
	productId := c.Param("id")
	var productBody models.Product

	if err := config.DB.Preload("Photo").Preload("Order").First(&productBody, productId).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Product not Found",
		})
		return
	}

	c.JSON(http.StatusOK, productBody)
}

func AddProduct(c *gin.Context) {
	var productInput struct {
		NamaDestinasi string `form:"namadestinasi"`
		Harga         int    `form:"harga"`
		Deskripsi     string `form:"deskripsi"`
		StartDate     string `form:"startdate"`
		EndDate       string `form:"enddate"`
		Thumbnail     string
	}

	err := c.ShouldBind(&productInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid entry",
		})
		return
	}

	if productInput.NamaDestinasi == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Destination Required.",
		})
		return
	}

	if productInput.Harga == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Price Required",
		})
		return
	}

	layout := "02/01/2006"
	parsedStartDate, err := time.Parse(layout, productInput.StartDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start date format. use dd/mm/yyyy"})
		return
	}

	parsedEndDate, err := time.Parse(layout, productInput.EndDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end date format. use dd/mm/yyyy"})
		return
	}

	// if productInput.Thumbnail == "" {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "Thumbnail Required",
	// 	})
	// }

	productBody := models.Product{
		NamaDestinasi: productInput.NamaDestinasi,
		Harga:         productInput.Harga,
		Deskripsi:     productInput.Deskripsi,
		StartDate:     parsedStartDate,
		EndDate:       parsedEndDate,
		Thumbnail:     productInput.Thumbnail,
	}

	err = config.DB.Create(&productBody).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to add Product",
		})
		return
	}

	thumbnail, err := c.FormFile("thumbnail")
	if err == nil {
		imageName := strings.ReplaceAll(thumbnail.Filename, " ", "_")
		imageFolder := fmt.Sprintf("./assets/image/thumbnail/%d", productBody.ID)
		err := os.MkdirAll(imageFolder, os.ModePerm)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create thumbnail folder.",
			})
			return
		}

		imagePath := fmt.Sprintf("%s/%s", imageFolder, imageName)
		err = c.SaveUploadedFile(thumbnail, imagePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to add thumbnail.",
			})
			return
		}
		productBody.Thumbnail = imageName
		err = config.DB.Save(&productBody).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to save product.",
			})
			return
		}
	}

	c.JSON(http.StatusOK, productBody)

}

func EditProduct(c *gin.Context) {
	productId := c.Param("id")

	var productData models.Product
	var productEditData struct {
		NamaDestinasi string `form:"edit-namadestinasi"`
		Harga         int    `form:"edit-harga"`
		Deskripsi     string `form:"edit-deskripsi"`
		StartDate     string `form:"edit-startdate"`
		EndDate       string `form:"edit-enddate"`
		Thumbnail     string
	}

	c.ShouldBind(&productEditData)

	err := config.DB.First(&productData, productId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Update failed. data not found.",
		})
		return
	}

	layout := "02/01/2006"

	updateData := make(map[string]interface{})

	if productEditData.NamaDestinasi != "" {
		updateData["NamaDestinasi"] = productEditData.NamaDestinasi
	}

	if productEditData.Harga != 0 {
		updateData["Harga"] = productEditData.Harga
	}

	if productEditData.Deskripsi != "" {
		updateData["Deskripsi"] = productEditData.Deskripsi
	}

	if productEditData.StartDate != "" {
		parsedStartDate, err := time.Parse(layout, productEditData.StartDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid start date format. use dd/mm/yyyy"})
			return
		}
		updateData["StartDate"] = parsedStartDate
	}

	if productEditData.EndDate != "" {
		parsedEndDate, err := time.Parse(layout, productEditData.EndDate)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid end date format. use dd/mm/yyyy"})
			return
		}
		updateData["EndDate"] = parsedEndDate
	}

	thumbnail, err := c.FormFile("edit-thumbnail")
	if err == nil {
		existingImgPath := fmt.Sprintf("./assets/image/thumbnail/%s/%s", productId, productData.Thumbnail)
		fmt.Println(existingImgPath)
		err := os.Remove(existingImgPath)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to update thumbnail",
			})
		}
		imageName := strings.ReplaceAll(thumbnail.Filename, " ", "_")
		imageFolder := fmt.Sprintf("./assets/image/thumbnail/%s", productId)
		err = os.MkdirAll(imageFolder, os.ModePerm)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to upload new Thumbnail.",
			})
			return
		}

		imagePath := fmt.Sprintf("%s/%s", imageFolder, imageName)
		err = c.SaveUploadedFile(thumbnail, imagePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to add thumbnail.",
			})
			return
		}

		updateData["Thumbnail"] = imageName

		// err = config.DB.Save(&updateData).Error
		// if err != nil {
		// 	c.JSON(http.StatusInternalServerError, gin.H{
		// 		"error": "failed to save product.",
		// 	})
		// 	return
		// }
	}

	if err = config.DB.Model(&productData).Updates(updateData).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to update product",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "update succesful.",
	})
}

func DeleteProduct(c *gin.Context) {
	var product models.Product
	var productPhoto models.ProductPhoto
	var orders models.Order
	productId, _ := strconv.Atoi(c.Param("id"))

	thumbnailPath := fmt.Sprintf("./assets/image/thumbnail/%d", productId)
	photosPath := fmt.Sprintf("./assets/image/productphoto/%d", productId)

	fmt.Println(thumbnailPath)
	err := os.RemoveAll(thumbnailPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Thumbnail deletion failed",
		})
		return
	}
	newErr := os.RemoveAll(photosPath)
	if newErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "photos deletion failed.",
		})
	}

	err = config.DB.Where("product_id=?", productId).Delete(&orders).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Product Photo Deletion Failed.",
		})
		return
	}

	err = config.DB.Where("product_id=?", productId).Delete(&productPhoto).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Product Photo Deletion Failed.",
		})
		return
	}

	err = config.DB.Where("ID=?", productId).Delete(&product).Error
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
