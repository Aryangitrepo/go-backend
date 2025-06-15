package routes

import (
	"intern/jwt"
	"intern/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login handles user authentication
func Login(service models.UserFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		var credentials models.LoginJson

		if err := c.ShouldBindJSON(&credentials); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "invalid login data",
				"details": err.Error(),
			})
			return
		}

		user, err := service.FindUser(&credentials)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid credentials",
			})
			return
		}

		token, err := jwt.CreateToken(user.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to generate token",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":  "Login successful",
			"token":    token,
			"user_id":  user.ID,
			"username": user.Name,
			"role":     user.Role,
		})
	}
}
