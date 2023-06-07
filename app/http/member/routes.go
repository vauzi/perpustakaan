package member

import (
	"github.com/gin-gonic/gin"
	"github.com/vauzi/perpustakaan/app/middleware"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func UserRoutes(public *gin.RouterGroup, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	r := public.Group("/members")
	r.Use(middleware.JwtAuthMiddleware())
	r.POST("/", middleware.UserIsActive, h.AddMembers)
	r.GET("/", h.GetAllMembers)
}
