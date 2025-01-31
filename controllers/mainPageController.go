package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func Katalog(c *gin.Context) {
	c.HTML(http.StatusOK, "katalog.html", nil)
}

func AboutUs(c *gin.Context) {
	c.HTML(http.StatusOK, "aboutus.html", nil)
}

func GalleryPage(c *gin.Context) {
	c.HTML(http.StatusOK, "gallery-main.html", nil)
}

func GalleryDetail(c *gin.Context) {
	c.HTML(http.StatusOK, "gallery-wisata.html", nil)
}

func DetailPage(c *gin.Context) {
	c.HTML(http.StatusOK, "detail.html", nil)
}

func OrderPage(c *gin.Context) {
	c.HTML(http.StatusOK, "booking.html", nil)
}

func Header(c *gin.Context) {
	c.File("./views/mainpage/header.html")
}

func Footer(c *gin.Context) {
	c.File("./views/mainpage/footer.html")
}

func Tatacara(c *gin.Context) {
	c.File("./views/mainpage/tatacara.html")
}

func Pilihan(c *gin.Context) {
	c.File("./views/mainpage/pilihan.html")
}
