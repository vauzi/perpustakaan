package member

import (
	"github.com/gin-gonic/gin"
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

	r.POST("/", h.AddMembers)
	r.GET("/", h.GetAllMembers)
}
