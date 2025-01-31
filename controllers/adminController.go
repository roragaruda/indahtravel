package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/attanabilrabbani/indahwisatabe/config"
	"github.com/attanabilrabbani/indahwisatabe/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func AdminCreate(c *gin.Context) {
	var adminBody models.Admin

	c.Bind(&adminBody)

	hash, err := bcrypt.GenerateFromPassword([]byte(adminBody.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Password Hashing failed",
		})
		return
	}

	admin := models.Admin{
		Username: adminBody.Username,
		Password: string(hash),
	}

	err = config.DB.Create(&admin).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error creating admin",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Succesful",
	})
}

func AdminLogin(c *gin.Context) {
	var adminLogin struct {
		Username string `form:"username"`
		Password string `form:"password"`
	}

	if c.ShouldBind(&adminLogin) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Login Failed",
		})
		return
	}

	var admin models.Admin

	config.DB.First(&admin, "username = ?", adminLogin.Username)
	if admin.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username or password.",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(adminLogin.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid username or password.",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"subj": admin.ID,
		"exp":  time.Now().Add(time.Hour).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create key",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Auth", tokenStr, 3600*10, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "login succesful.",
	})

}

func AdminLogout(c *gin.Context) {
	c.SetCookie("Auth", "", -1, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout Succesful.",
	})
}

func AdminValidate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"valid": true})
}
