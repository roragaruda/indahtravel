package routes

import (
	"github.com/attanabilrabbani/indahwisatabe/controllers"
	"github.com/gin-gonic/gin"
)

func MainPageRoutes(r *gin.Engine) {
	indahwisata := r.Group("/indahwisata") // Prefix "indahwisata"
	{
		indahwisata.GET("/", controllers.Index)
		indahwisata.GET("/katalog", controllers.Katalog)
		indahwisata.GET("/gallery", controllers.GalleryPage)
		indahwisata.GET("/aboutus", controllers.AboutUs)
		indahwisata.GET("/unauthorized", controllers.Unauthorized)
		indahwisata.GET("/detail/:id", controllers.DetailPage)
		indahwisata.GET("/gallery/:id", controllers.GalleryDetail)
		indahwisata.GET("/order/:id", controllers.OrderPage)
	}

	r.GET("/header.html", controllers.Header)
	r.GET("/footer.html", controllers.Footer)
	r.GET("/tatacara.html", controllers.Tatacara)
	r.GET("/pilihan.html", controllers.Pilihan)
}
