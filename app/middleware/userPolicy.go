package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/vauzi/perpustakaan/app/models"
	"github.com/vauzi/perpustakaan/app/token"
)

func UserBorrwerPolicy(c *gin.Context) {
	idString := c.Param("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid ID format"})
		return
	}

	db := models.ConnectDatabase()

	var borrower = models.Borrower{}

	if result := db.First(&borrower, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	userID, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if userID != borrower.UserID {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "error", "message": "you do not have access to this"})
		return
	}
	c.Next()
}
