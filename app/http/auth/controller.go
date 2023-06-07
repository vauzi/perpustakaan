package auth

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/vauzi/perpustakaan/app/models"
	"golang.org/x/crypto/bcrypt"
)

type AddUserRequestBody struct {
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=6"`
}

var jwtSecret = []byte("your_secret_key")

func (h handler) SignUp(c *gin.Context) {
	body := AddUserRequestBody{}

	// hendler validation
	if err := c.ShouldBindJSON(&body); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) {
			out := make([]ErrorMsg, len(ve))
			for i, fe := range ve {
				out[i] = ErrorMsg{fe.Field(), getErrorMsg(fe)}
			}
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"errors": out})
		}
		return
	}

	var user = models.User{}

	result := h.DB.First(&user, "username = ?", body.Username)
	if result.Error == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": gin.H{"field": "Username", "message": "This field is unique"}})
		return
	}
	user.Username = body.Username
	user.Password = body.Password
	user.HashPassword()

	if result := h.DB.Create(&user); result.Error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"errors": result.Error})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created"})
}

func (h handler) SignIp(c *gin.Context) {
	body := AddUserRequestBody{}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user = models.User{}
	if result := h.DB.First(&user, "username = ?", body.Username); result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	// Create JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		log.Println("Failed to generate token:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Signed in successfully", "token": tokenString})

}
