package category

import (
	"github.com/gin-gonic/gin"
	"github.com/vauzi/perpustakaan/app/middleware"
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
	r.Use(middleware.JwtAuthMiddleware())
	r.POST("/", middleware.UserIsActive, h.Createcategory)
	r.GET("/", h.GetCategory)
}
