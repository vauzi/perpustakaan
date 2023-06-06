package book

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/vauzi/perpustakaan/app/models"
)

type AddBooksRequestBody struct {
	Title       string    `json:"title" binding:"required"`
	Author      string    `json:"author" binding:"required"`
	Published   string    `json:"published" binding:"required"`
	Description string    `json:"description" binding:"required"`
	CategoryID  uuid.UUID `json:"category_id" binding:"required"`
}

type BookResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Published   string    `json:"published"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
}

func (h handler) AddBooks(c *gin.Context) {
	var body = AddBooksRequestBody{}

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

	var book = models.Book{}
	book.Title = body.Title
	book.Author = body.Author
	book.Published = body.Published
	book.Description = body.Description
	book.CategoryID = body.CategoryID

	if result := h.DB.Create(&book); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": result.Error})
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "books added successfully"})
}

func (h handler) GetAllBooks(c *gin.Context) {

	var books []models.Book

	if result := h.DB.Preload("Category").Find(&books); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": result.Error})
		return
	}

	var bookResponses []BookResponse
	for _, book := range books {
		bookResponses = append(bookResponses, BookResponse{
			ID:          book.ID,
			Title:       book.Title,
			Author:      book.Author,
			Published:   book.Published,
			Description: book.Description,
			Category:    book.Category.Name,
		})
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "data": bookResponses})
}
