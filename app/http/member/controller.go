package member

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/vauzi/perpustakaan/app/models"
)

type MemberRequestBody struct {
	FullName    string    `json:"full_name" binding:"required"`
	NoInduk     string    `json:"no_induk" binding:"required"`
	NoHp        string    `json:"no_hp" binding:"required"`
	Gender      string    `json:"gender" binding:"required"`
	Work        string    `json:"work" binding:"required"`
	UserAddress string    `json:"user_address" binding:"required"`
	Status      string    `json:"status" binding:"required"`
	BirthDate   time.Time `json:"birth_date" binding:"required"`
	BirthPlace  string    `json:"birth_place" binding:"required"`
}

func (h handler) AddMembers(c *gin.Context) {
	var body = MemberRequestBody{}

	if err := c.ShouldBindJSON(&body); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "errors", "message": out})
		}
		return
	}

	var member = models.Member{}
	member.FullName = body.FullName
	member.NoInduk = body.NoInduk
	member.NoHp = body.NoHp
	member.Gender = body.Gender
	member.Work = body.Work
	member.UserAddress = body.UserAddress
	member.Status = body.Status
	member.BirthDate = body.BirthDate
	member.BirthPlace = body.BirthPlace

	if result := h.DB.Create(&member); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": result.Error})
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "user added successfully"})

}

func (h handler) GetAllMembers(c *gin.Context) {

	var member []models.Member

	if result := h.DB.Find(&member); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": result.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": member})
}
