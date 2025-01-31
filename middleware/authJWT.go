package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/attanabilrabbani/indahwisatabe/config"
	"github.com/attanabilrabbani/indahwisatabe/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequireAuth(c *gin.Context) {
	tokenStr, err := c.Cookie("Auth")

	if err != nil {
		c.Redirect(http.StatusFound, "/indahwisata/unauthorized")
		// c.JSON(http.StatusUnauthorized, gin.H{"valid": false})
		c.Abort()
		return
	}

	token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//check for exp date
		if float64(time.Now().Unix()) > claims["exp"].(float64) {
			// c.JSON(http.StatusUnauthorized, gin.H{"valid": false})
			c.Redirect(http.StatusFound, "/indahwisata/unauthorized")
			c.Abort()
			return
		}

		//find admin
		var admin models.Admin
		config.DB.First(&admin, "ID = ?", claims["subj"])
		if admin.ID == 0 {
			// c.JSON(http.StatusUnauthorized, gin.H{
			// 	"valid": false,
			// })
			c.Redirect(http.StatusFound, "/indahwisata/unauthorized")
			c.Abort()
			return
		}

		c.Set("admin", admin)
		c.Next()
	} else {
		// c.JSON(http.StatusUnauthorized, gin.H{"valid": false})
		c.Redirect(http.StatusFound, "/indahwisata/unauthorized")
		c.Abort()
		return
	}
}
