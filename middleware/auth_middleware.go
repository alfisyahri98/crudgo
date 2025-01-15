package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"projectgo/utils"
)

func MiddlewareToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		accessToken, err := c.Cookie("access_token")
		if err != nil || accessToken == "" {
			refreshToken, err := c.Cookie("refresh_token")
			if err != nil || refreshToken == "" {
				c.JSON(http.StatusUnauthorized, utils.Response{
					Status:  http.StatusUnauthorized,
					Message: "Unauthorized",
				})
				c.Abort()
				return
			}

			userID, err := utils.ValidateRefreshToken(refreshToken)
			if err != nil {
				c.JSON(http.StatusUnauthorized, utils.Response{
					Status:  http.StatusUnauthorized,
					Message: "Refresh Token Invalid",
				})
				c.Abort()
				return
			}

			newAccessToken, err := utils.GenerateJWTAccessToken(userID)
			if err != nil {
				c.JSON(http.StatusUnauthorized, utils.Response{
					Status:  http.StatusUnauthorized,
					Message: "Failed Generate New Access Token",
				})
				c.Abort()
				return
			}

			newRefreshToken, err := utils.GenerateJWTRefreshToken(userID)
			if err != nil {
				c.JSON(http.StatusUnauthorized, utils.Response{
					Status:  http.StatusUnauthorized,
					Message: "Failed Generate New Refresh Token",
				})
				c.Abort()
				return
			}

			c.SetCookie("access_token", newAccessToken, 10, "", "", false, true)
			c.SetCookie("refresh_token", newRefreshToken, 3600, "", "", false, true)
			c.Next()
			return
		}

		_, err = utils.ValidateAccessToken(accessToken)
		if err != nil {
			c.JSON(http.StatusUnauthorized, utils.Response{
				Status:  http.StatusUnauthorized,
				Message: "Access Token Invalid",
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
		c.SetCookie("csrf_token", csrfToken, 200, "/", "", false, true)
		c.Next()
	}
}
