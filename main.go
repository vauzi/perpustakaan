package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/vauzi/perpustakaan/app/http/auth"
	"github.com/vauzi/perpustakaan/app/http/book"
	"github.com/vauzi/perpustakaan/app/http/category"
	"github.com/vauzi/perpustakaan/app/http/member"
	"github.com/vauzi/perpustakaan/app/models"
)

func main() {
	r := gin.Default()

	public := r.Group("/api/v1")
	dbHandler := models.ConnectDatabase()

	auth.AuthRoutes(public, dbHandler)
	category.CategoryRoutes(public, dbHandler)
	book.BookRoutes(public, dbHandler)
	member.UserRoutes(public, dbHandler)

	if err := godotenv.Load(); err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("SERVER_PORT")

	r.Run(port)

}
