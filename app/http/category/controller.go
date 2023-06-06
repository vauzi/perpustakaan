package category

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/vauzi/perpustakaan/app/models"
)

type AddCategoryRequestBody struct {
	Name string `json:"name" binding:"required"`
}

func (h handler) Createcategory(c *gin.Context) {
	body := AddCategoryRequestBody{}

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

	var category = models.Category{}
	category.Name = body.Name

	if result := h.DB.Create(&category); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": result.Error})
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "create category successfully"})
}
