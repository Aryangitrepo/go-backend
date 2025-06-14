package middleware

import (
	"intern/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	reqToken := c.GetHeader("Authorization")

	if err := jwt.VerifyToken(reqToken); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"ERROR": "INCORRECT TOKEN .." + err.Error() + reqToken,
		})
		c.Abort()
	} else {
		c.Next()
	}

}
