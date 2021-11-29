package controller

import (
	"Abinash393/digital-trons-test/model"
	"Abinash393/digital-trons-test/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var currentUser model.User
	c.BindJSON(&currentUser)

	var user model.User
	model.DB.Where(model.User{Email: currentUser.Email}).First(&user)

	if user.CheckPassword(currentUser.Password) {
		token := util.GenerateToken(&user)
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"token":   token,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "invalid credentials",
		})
	}
}
