package routes

import (
	"github.com/attanabilrabbani/indahwisatabe/controllers"
	"github.com/attanabilrabbani/indahwisatabe/middleware"
	"github.com/gin-gonic/gin"
)

func APIRoutes(r *gin.Engine) {
	r.GET("/product", controllers.GetProduct)
	r.GET("/product/:id", controllers.GetProductById)
	r.GET("/orders", controllers.GetOrder)
	r.GET("/orders/search/:code", controllers.GetOrderbyCode)
	r.GET("/orders/:id", controllers.GetOrderById)
	r.GET("/gallery", controllers.GetGallery)
	r.GET("/gallery/:id", controllers.GetGalleryById)
	r.POST("/product/add", middleware.RequireAuth, controllers.AddProduct)
	r.POST("/orders/add", controllers.CreateOrder)
	r.POST("/product/photos/:id", middleware.RequireAuth, controllers.UploadProductPhoto)
	r.POST("/gallery/add", middleware.RequireAuth, controllers.CreateGallery)
	r.POST("/gallery/photos/add/:id", middleware.RequireAuth, controllers.UploadGalleryPhotos)
	r.PUT("/product/edit/:id", middleware.RequireAuth, controllers.EditProduct)
	r.PUT("/gallery/edit/:id", middleware.RequireAuth, controllers.EditGalleryName)
	r.PUT("/orders/edit/:id", middleware.RequireAuth, controllers.UpdateOrder)
	r.DELETE("/orders/:id", middleware.RequireAuth, controllers.DeleteOrder)
	r.DELETE("/product/delete/:id", middleware.RequireAuth, controllers.DeleteProduct)
	r.DELETE("/product/photos/:id", middleware.RequireAuth, controllers.DeletePhoto)
	r.DELETE("/gallery/:id", middleware.RequireAuth, controllers.DeleteGallery)
	r.DELETE("/gallery/photos/:id", middleware.RequireAuth, controllers.DeleteGalleryPhotos)
}
