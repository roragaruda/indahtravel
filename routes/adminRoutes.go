package routes

import (
	"github.com/attanabilrabbani/indahwisatabe/controllers"
	"github.com/attanabilrabbani/indahwisatabe/middleware"
	"github.com/gin-gonic/gin"
)

func AdminPageRoutes(r *gin.Engine) {
	r.GET("/indahadmin/adminlogin", controllers.Login)
	r.GET("/index", middleware.RequireAuth, controllers.AdminHome)
	r.GET("/validate", middleware.RequireAuth, controllers.AdminValidate)
	r.GET("/product", middleware.RequireAuth, controllers.ProductPage)
	r.GET("/orders", middleware.RequireAuth, controllers.OrdersPage)
	r.GET("/productphoto", middleware.RequireAuth, controllers.ProductPhoto)
	r.GET("/gallery", middleware.RequireAuth, controllers.Gallery)
	r.GET("/menu.html", controllers.Menu)
	r.POST("/create", controllers.AdminCreate)
	r.POST("/login", controllers.AdminLogin)
	r.POST("/logout", middleware.RequireAuth, controllers.AdminLogout)
	r.POST("/admin/add", controllers.AdminCreate)
}
