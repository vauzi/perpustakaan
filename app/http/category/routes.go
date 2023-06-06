package category

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func CategoryRoutes(public *gin.RouterGroup, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	r := public.Group("/category")
	r.POST("/", h.Createcategory)
}
