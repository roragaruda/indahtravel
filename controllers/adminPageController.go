package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func AdminHome(c *gin.Context) {
	c.HTML(http.StatusOK, "homepage.html", nil)
}

func ProductPage(c *gin.Context) {
	c.HTML(http.StatusOK, "product.html", nil)
}

func OrdersPage(c *gin.Context) {
	c.HTML(http.StatusOK, "orders.html", nil)
}

func ProductPhoto(c *gin.Context) {
	c.HTML(http.StatusOK, "productphoto.html", nil)
}

func Gallery(c *gin.Context) {
	c.HTML(http.StatusOK, "gallery.html", nil)
}

func Unauthorized(c *gin.Context) {
	c.HTML(http.StatusUnauthorized, "unauthorized.html", nil)
}

func Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func Menu(c *gin.Context) {
	c.File("./views/admin/menu.html")
}
