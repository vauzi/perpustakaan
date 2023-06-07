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
	UserID       uuid.UUID  `gorm:"not null" json:"user_id"`
	Member       Member     `gorm:"foreignKey:UserID" json:"members"`
	AdminID      uuid.UUID  `gorm:"not null" json:"admin_id"`
	Admin        Admin      `gorm:"foreignKey:UserID" json:"admins"`
	BookID       uuid.UUID  `gorm:"not null" json:"book_id"`
	Book         Book       `gorm:"foreignKey:UserID" json:"books"`
}
