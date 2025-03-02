package routes

import (
	"github.com/attanabilrabbani/indahwisatabe/controllers"
	"github.com/gin-gonic/gin"
)

func MainPageRoutes(r *gin.Engine) {
	r.GET("/", controllers.Index)
	r.GET("/katalog", controllers.Katalog)
	r.GET("/gallery", controllers.GalleryPage)
	r.GET("/aboutus", controllers.AboutUs)
	r.GET("/unauthorized", controllers.Unauthorized)
	r.GET("/detail/:id", controllers.DetailPage)
	r.GET("/gallery/:id", controllers.GalleryDetail)
	r.GET("/order/:id", controllers.OrderPage)
	r.GET("/header.html", controllers.Header)
	r.GET("/footer.html", controllers.Footer)
	r.GET("/tatacara.html", controllers.Tatacara)
	r.GET("/pilihan.html", controllers.Pilihan)
}
