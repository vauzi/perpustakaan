package borrower

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/vauzi/perpustakaan/app/models"
	"github.com/vauzi/perpustakaan/app/token"
)

type BorroerRequestBody struct {
	MemberID uuid.UUID `json:"member_id" binding:"required"`
	BookID   uuid.UUID `json:"book_id" binding:"required"`
}

func (h handler) AddBorrowerBook(c *gin.Context) {
	var body = BorroerRequestBody{}

	if err := c.ShouldBindJSON(&body); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "errors", "message": "out"})
		}
		return
	}

	userID, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var borrower = models.Borrower{}
	borrower.BorrowedDate = time.Now()
	borrower.MemberID = body.MemberID
	borrower.UserID = userID
	borrower.BookID = body.BookID

	if result := h.DB.Create(&borrower); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": userID, "message": result.Error})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "create successfully"})
}

func (h handler) GetAllBorrowers(c *gin.Context) {
	var borrower []models.Borrower

	if result := h.DB.Find(&borrower); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": borrower})
}

func (h handler) DeleteBorrowers(c *gin.Context) {
	idString := c.Param("id")
	id, err := uuid.Parse(idString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid ID format"})
		return
	}

	var borrower = models.Borrower{}

	if result := h.DB.First(&borrower, id); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": result.Error.Error()})
		return
	}

	h.DB.Delete(&borrower, id)
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "delete borrowed successfully"})

}
