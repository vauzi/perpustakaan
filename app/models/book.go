package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	ID          uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Title       string     `gorm:"size:255;not null" json:"title"`
	Author      string     `gorm:"size:255;not null" json:"author"`
	Published   string     `gorm:"size:255;not null" json:"published"`
	Description string     `gorm:"type:text" json:"description"`
	CategoryID  uuid.UUID  `gorm:"not null" json:"category_id"`
	Category    Category   `gorm:"foreignKey:CategoryID" json:"category"`
	Borrowers   []Borrower `gorm:"foreignKey:BookID" json:"borrowers"`
}
