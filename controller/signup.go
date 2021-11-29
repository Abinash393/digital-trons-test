package controller

import (
	"Abinash393/digital-trons-test/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Signup(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)

	model.DB.Create(&user)
	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "user created",
	})
}
