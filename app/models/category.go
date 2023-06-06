package models

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	ID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name string    `gorm:"size:255;not null" json:"name"`
}
