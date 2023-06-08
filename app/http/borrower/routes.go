package borrower

import (
	"github.com/gin-gonic/gin"
	"github.com/vauzi/perpustakaan/app/middleware"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func BorrowerRoutes(public *gin.RouterGroup, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	r := public.Group("/borrower")
	r.Use(middleware.JwtAuthMiddleware())

	r.POST("/", h.AddBorrowerBook)
	r.GET("/", h.GetAllBorrowers)
	r.DELETE("/:id", middleware.UserBorrwerPolicy, h.DeleteBorrowers)
}
