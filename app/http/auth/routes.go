package auth

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func AuthRoutes(public *gin.RouterGroup, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	public.POST("/sign-up", h.SignUp)
	public.POST("/sign-ip", h.SignIp)
}
