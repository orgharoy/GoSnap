package model

import (
	"time"

	"github.com/google/uuid"
)

type Post struct {
	ID          uuid.UUID `gorm:"type:uuid; default:uuid_generate_v4(); primary_key" json:"id"`
	Title       string    `gorm:"varchar(255); not null" json:"title"`
	Description string    `gorm:"varchar(255); not null" json:"description"`
	Image       string    `gorm:"varchar(255)" json:"image"`
	Address     string    `gorm:"null" json:"address"`
	CreatedAt   time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"not null" json:"updatedAt"`
	UserID      uuid.UUID `gorm:"type:uuid; not null" json:"userId"`
}
