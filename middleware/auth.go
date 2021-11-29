package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IsAuthorized(c *gin.Context) {
	if c.Request.Header["Token"] == nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"message": "Token not found",
		})
	}
}
