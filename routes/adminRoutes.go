package routes

import (
	"github.com/attanabilrabbani/indahwisatabe/controllers"
	"github.com/attanabilrabbani/indahwisatabe/middleware"
	"github.com/gin-gonic/gin"
)

func AdminPageRoutes(r *gin.Engine) {
	r.GET("/indahadmin/adminlogin", controllers.Login)
	r.GET("/indahadmin/index", middleware.RequireAuth, controllers.AdminHome)
	r.GET("/indahadmin/validate", middleware.RequireAuth, controllers.AdminValidate)
	r.GET("/indahadmin/product", middleware.RequireAuth, controllers.ProductPage)
	r.GET("/indahadmin/orders", middleware.RequireAuth, controllers.OrdersPage)
	r.GET("/indahadmin/productphoto", middleware.RequireAuth, controllers.ProductPhoto)
	r.GET("/indahadmin/gallery", middleware.RequireAuth, controllers.Gallery)
	r.GET("/menu.html", controllers.Menu)
	r.POST("/indahadmin/create", controllers.AdminCreate)
	r.POST("/indahadmin/login", controllers.AdminLogin)
	r.POST("/indahadmin/logout", middleware.RequireAuth, controllers.AdminLogout)
	r.POST("/admin/add", controllers.AdminCreate)
}
