package models

import (
	"html"
	"strings"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Username string    `gorm:"size:255;not null;unique" json:"username"`
	Password string    `gorm:"size:255;not null;" json:"password"`
}

func (user *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	user.Username = html.EscapeString(strings.TrimSpace(user.Username))

	return nil
}
