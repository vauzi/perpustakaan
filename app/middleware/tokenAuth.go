package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vauzi/perpustakaan/app/models"
	"github.com/vauzi/perpustakaan/app/token"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Unautorize"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func UserIsActive(c *gin.Context) {

	userID, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	db := models.ConnectDatabase()

	var user models.User
	if result := db.First(&user, userID); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if user.IsActive != true {
		c.AbortWithStatusJSON(http.StatusExpectationFailed, gin.H{"error": "Your account is not active"})
		return
	}

	c.Next()
}
