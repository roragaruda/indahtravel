package controllers

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/attanabilrabbani/indahwisatabe/config"
	"github.com/attanabilrabbani/indahwisatabe/models"
	"github.com/gin-gonic/gin"
)

func CreateOrder(c *gin.Context) {
	var orderInput models.Order
	// var orderInput struct {
	// 	NamaPeserta   string `form:"nama-peserta"`
	// 	NoTelp        string `form:"no-telp"`
	// 	JmlPeserta    uint   `form:"jml-peserta"`
	// 	FotoBukti     string
	// 	KodePembelian string `form:"kode-pembelian"`
	// 	Status        string `form:"status"`
	// 	ProductID     uint   `form:"product-id"`
	// }

	err := c.ShouldBind(&orderInput)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Order",
		})
		return
	}

	if orderInput.NamaPeserta == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name Required",
		})
		return
	}

	if orderInput.NoTelp == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Phone Required",
		})
		return
	}

	if orderInput.JmlPeserta == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Amount of People required.",
		})
		return
	}

	if orderInput.ProductID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Product ID Required.",
		})
		return
	}

	if orderInput.KodePembelian == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "missing order code.",
		})
		return
	}

	err = config.DB.Create(&orderInput).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Order Placement Failed",
		})
		return
	}

	transaction, err := c.FormFile("foto-bukti")
	if err == nil {
		imageName := strings.ReplaceAll(transaction.Filename, " ", "_")
		imageFolder := fmt.Sprintf("./assets/image/bukti/%d", orderInput.ID)
		err := os.MkdirAll(imageFolder, os.ModePerm)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create folder.",
			})
			return
		}

		imagePath := fmt.Sprintf("%s/%s", imageFolder, imageName)
		err = c.SaveUploadedFile(transaction, imagePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to add order.",
			})
			return
		}
		orderInput.FotoBukti = imageName
		err = config.DB.Save(&orderInput).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to save order.",
			})
			return
		}
	}

	c.JSON(http.StatusOK, orderInput)

}

func GetOrder(c *gin.Context) {
	var order []models.Order

	err := config.DB.Find(&order).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Unable to fetch order data.",
		})
		return
	}

	c.JSON(http.StatusOK, order)

}

func GetOrderById(c *gin.Context) {
	var order models.Order
	orderId := c.Param("id")

	err := config.DB.Where("ID=?", orderId).First(&order).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Order not found.",
		})
		return
	}

	c.JSON(http.StatusOK, order)
}

func GetOrderbyCode(c *gin.Context) {
	code := c.Param("code")
	var order models.Order

	err := config.DB.Where("kode_pembelian=?", code).First(&order).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Order not found.",
		})
		return
	}

	c.JSON(http.StatusOK, order)

}

func UpdateOrder(c *gin.Context) {
	var order models.Order
	var orderEdit struct {
		NamaPeserta   string `form:"nama-peserta"`
		NoTelp        string `form:"no-telp"`
		JmlPeserta    uint   `form:"jml-peserta"`
		FotoBukti     string
		KodePembelian string `form:"kode-pembelian"`
		Status        string `form:"status"`
		ProductID     uint   `form:"product-id"`
	}
	orderId := c.Param("id")

	c.ShouldBind(&orderEdit)

	err := config.DB.First(&order, orderId).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Order not found.",
		})
		return
	}

	orderEditData := make(map[string]interface{})

	if orderEdit.NamaPeserta != "" {
		orderEditData["NamaPeserta"] = orderEdit.NamaPeserta
	}
	if orderEdit.NoTelp != "" {
		orderEditData["NoTelp"] = orderEdit.NoTelp
	}
	if orderEdit.JmlPeserta != 0 {
		orderEditData["JmlPeserta"] = orderEdit.JmlPeserta
	}
	if orderEdit.Status != "" {
		orderEditData["Status"] = orderEdit.Status
	}
	if orderEdit.ProductID != 0 {
		orderEditData["ProductID"] = orderEdit.ProductID
	}

	err = config.DB.Model(&order).Updates(orderEditData).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Update Failed",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Update Successful",
	})

}

func DeleteOrder(c *gin.Context) {
	var order models.Order
	orderId := c.Param("id")

	err := config.DB.Where("ID=?", orderId).Delete(&order).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Deletion Failed.",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Order deleted successfully",
	})
}
