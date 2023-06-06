package book

import (
	"github.com/gin-gonic/gin"
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

	r.POST("/", h.AddBooks)
	r.GET("/", h.GetAllBooks)
}
