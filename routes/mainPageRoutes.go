package routes

import (
	"github.com/attanabilrabbani/indahwisatabe/controllers"
	"github.com/gin-gonic/gin"
)

func MainPageRoutes(r *gin.Engine) {
	routes := r.Group("/indahwisata")
	{
		routes.GET("/", controllers.Index)
		routes.GET("/katalog", controllers.Katalog)
		routes.GET("/gallery", controllers.GalleryPage)
		routes.GET("/aboutus", controllers.AboutUs)
		routes.GET("/unauthorized", controllers.Unauthorized)
		routes.GET("/detail/:id", controllers.DetailPage)
		routes.GET("/gallery/:id", controllers.GalleryDetail)
		routes.GET("/order/:id", controllers.OrderPage)
		routes.GET("/header.html", controllers.Header)
		routes.GET("/footer.html", controllers.Footer)
		routes.GET("/tatacara.html", controllers.Tatacara)
		routes.GET("/pilihan.html", controllers.Pilihan)
	}
}
