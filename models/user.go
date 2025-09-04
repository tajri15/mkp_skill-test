package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	FullName     string    `gorm:"size:100;not null" json:"full_name"`
	Email        string    `gorm:"size:100;not null;unique" json:"email"`
	PasswordHash string    `gorm:"size:255;not null" json:"-"`
	Role         string    `gorm:"size:20;not null" json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

// Hook untuk hash password sebelum menyimpan user (jika ada fitur registrasi)
func (user *User) BeforeSave(tx *gorm.DB) (err error) {
	if user.PasswordHash != "" {
		// Cek apakah password sudah di-hash atau belum
		// Jika password belum di-hash, maka hash
		if _, err := bcrypt.Cost([]byte(user.PasswordHash)); err != nil {
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
			if err != nil {
				return err
			}
			user.PasswordHash = string(hashedPassword)
		}
	}
	return
}