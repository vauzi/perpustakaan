package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Borrower struct {
	gorm.Model
	ID           uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	BorrowedDate time.Time  `gorm:"type:timestamp;default:CURRENT_TIMESTAMP" json:"borrowed_date"`
	ReturnedDate *time.Time `gorm:"type:timestamp;default:null" json:"returned_date"`
	UserID       uint       `gorm:"not null" json:"user_id"`
	User         User       `gorm:"foreignKey:UserID" json:"users"`
}
