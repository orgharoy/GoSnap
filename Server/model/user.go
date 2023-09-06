package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID             uuid.UUID `gorm:"type:uuid; default:uuid_generate_v4(); primary_key" json:"id"`
	FirstName      string    `gorm:"varchar(255); not null" json:"firstName"`
	LastName       string    `gorm:"varchar(255); not null" json:"lastName"`
	Email          string    `gorm:"varchar(255); not null" json:"email"`
	ProfilePicture string    `gorm:"varchar(255)" json:"profilePicture"`
	Bio            string    `gorm:"null" json:"bio"`
	Address        string    `gorm:"null" json:"address"`
	Password       string    `gorm:"not null" json:"password"`
	CreatedAt      time.Time `gorm:"not null" json:"createdAt"`
	UpdatedAt      time.Time `gorm:"not null" json:"updatedAt"`
}
