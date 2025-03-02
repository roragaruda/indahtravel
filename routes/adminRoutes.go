package routes

import (
	"github.com/attanabilrabbani/indahwisatabe/controllers"
	"github.com/attanabilrabbani/indahwisatabe/middleware"
	"github.com/gin-gonic/gin"
)

func AdminPageRoutes(r *gin.Engine) {
	admins := r.Group("/indahadmin")
	{
		admins.GET("/adminlogin", controllers.Login)
		admins.GET("/index", middleware.RequireAuth, controllers.AdminHome)
		admins.GET("/validate", middleware.RequireAuth, controllers.AdminValidate)
		admins.GET("/product", middleware.RequireAuth, controllers.ProductPage)
		admins.GET("/orders", middleware.RequireAuth, controllers.OrdersPage)
		admins.GET("/productphoto", middleware.RequireAuth, controllers.ProductPhoto)
		admins.GET("/gallery", middleware.RequireAuth, controllers.Gallery)
		admins.GET("/menu.html", controllers.Menu)
		admins.POST("/create", controllers.AdminCreate)
		admins.POST("/login", controllers.AdminLogin)
		admins.POST("/logout", middleware.RequireAuth, controllers.AdminLogout)
		admins.POST("/admin/add", controllers.AdminCreate)
	}
}
