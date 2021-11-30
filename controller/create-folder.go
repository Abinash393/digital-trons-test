package controller

import (
	"Abinash393/digital-trons-test/db"
	"Abinash393/digital-trons-test/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateFolder(c *gin.Context) {
	currentUserID, exist := c.Get("currentUserID")
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "user id not found in token",
		})
	} else {
		id := uint(currentUserID.(float64))
		var currentUser model.User
		db.ORM.Where(model.User{ID: id}, &currentUser)

		folder := model.Folder{}
		db.ORM.Create(&folder)
		c.JSON(http.StatusOK, gin.H{
			"success":       true,
			"folderCreated": folder,
		})
	}
}
