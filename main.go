package main

import (
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

	r.Run()
}
