package book

import (
	"github.com/gin-gonic/gin"
	"github.com/vauzi/perpustakaan/app/middleware"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func BookRoutes(public *gin.RouterGroup, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	r := public.Group("/books")
	r.Use(middleware.JwtAuthMiddleware())

	r.POST("/", middleware.UserIsActive, h.AddBooks)
	r.GET("/", h.GetAllBooks)
}
