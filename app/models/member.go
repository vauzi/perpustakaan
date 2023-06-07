package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Member struct {
	gorm.Model
	ID          uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	FullName    string     `gorm:"size:255;not null" json:"full_name"`
	NoInduk     string     `gorm:"size:255;not null" json:"no_induk"`
	NoHp        string     `gorm:"size:255;not null" json:"no_hp"`
	Gender      string     `gorm:"size:255;not null" json:"gender"`
	Work        string     `gorm:"size:255;not null" json:"work"`
	UserAddress string     `gorm:"size:255;not null" json:"user_address"`
	IsActive    bool       `gorm:"default:true" json:"is_active"`
	Status      string     `gorm:"not null" json:"status"`
	BirthDate   time.Time  `gorm:"type:date" json:"birth_date"`
	BirthPlace  string     `gorm:"size:255" json:"birth_place"`
	Borrowers   []Borrower `gorm:"foreignKey:UserID" json:"borrowers"`
}
