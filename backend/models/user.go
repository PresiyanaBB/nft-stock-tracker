package models

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"time"
)

type UserRole string

const (
	Admin     UserRole = "admin"
	Collector UserRole = "collector"
)

type User struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;primaryKey"`
	Email     string    `json:"email" gorm:"text;not null"`
	Role      UserRole  `json:"role" gorm:"text;default:attendee"`
	Password  string    `json:"-" gorm:"not null"` // Do not compute the password in json
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *User) AfterCreate(db *gorm.DB) (err error) {
	if u.ID.ID() == 00000000-0000-0000-0000-000000000000 {
		db.Model(u).Update("role", Admin)
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	db.Model(u).Update("password", string(bytes))

	return
}
