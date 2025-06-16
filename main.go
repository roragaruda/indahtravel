package main

import (
	"os"
	"github.com/attanabilrabbani/indahwisatabe/config"
	"github.com/attanabilrabbani/indahwisatabe/routes"
	"github.com/gin-gonic/gin"
)

func init() {
	config.LoadEnvVars()
	config.ConnectDB()
	config.MigrateDB()
}

func main() {
	r := gin.Default()

	r.Static("/static", "./views/static")
	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("views/**/*.html")

	routes.APIRoutes(r)
	routes.AdminPageRoutes(r)
	routes.MainPageRoutes(r)

	// Redirect root URL ke /indahwisata
	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/indahwisata")
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
