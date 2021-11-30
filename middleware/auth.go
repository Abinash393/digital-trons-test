package middleware

import (
	"Abinash393/digital-trons-test/util"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsAuthorized(c *gin.Context) {
	if c.Request.Header["Authorization"] == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Token not found",
		})
	}

	uid := util.DecodeToken(c.Request.Header["Authorization"][0])
	fmt.Println("\n", "UID:", uid)
	if uid == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Invalid token",
		})
		return
	} else {
		c.Set("currentUserID", uid)
		c.Next()
	}
}
