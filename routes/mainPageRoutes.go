package routes

import (
	"github.com/attanabilrabbani/indahwisatabe/controllers"
	"github.com/gin-gonic/gin"
)

func MainPageRoutes(r *gin.Engine) {
	// Redirect dari root ke /indahwisata
	r.GET("/", func(c *gin.Context) {
		c.Redirect(302, "/indahwisata")
	})

	r.GET("/indahwisata", controllers.Index)
	r.GET("/indahwisata/katalog", controllers.Katalog)
	r.GET("/indahwisata/gallery", controllers.GalleryPage)
	r.GET("/indahwisata/aboutus", controllers.AboutUs)
	r.GET("/indahwisata/unauthorized", controllers.Unauthorized)
	r.GET("/indahwisata/detail/:id", controllers.DetailPage)
	r.GET("/indahwisata/gallery/:id", controllers.GalleryDetail)
	r.GET("/indahwisata/order/:id", controllers.OrderPage)
	r.GET("/header.html", controllers.Header)
	r.GET("/footer.html", controllers.Footer)
	r.GET("/tatacara.html", controllers.Tatacara)
	r.GET("/pilihan.html", controllers.Pilihan)
}
