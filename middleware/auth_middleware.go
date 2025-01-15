package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectgo/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, _ := c.Cookie("access_token")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": http.StatusUnauthorized,
			})
			c.Abort()
			return
		}

		_, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token",
			})
			c.Abort()
			return
		}

		c.Next()
	}
}

func CSRFMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		csrfToken, err := utils.GenerateCSRFToken()
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Failed to generate csrf token",
			})
			c.Abort()
			return
		}
		c.SetCookie("csrf_token", csrfToken, 3600, "/", "", false, true)
		c.Next()
	}
}
