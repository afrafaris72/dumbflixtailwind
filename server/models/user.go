package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Fullname  string    `json:"fullName" gorm:"type: VARCHAR(255)"`
	Email     string    `json:"email" binding:"required, email" gorm:"unique; not null"`
	Password  string    `json:"password" gorm:"type: VARCHAR(255)"`
	Gender    string    `json:"gender" gorm:"type: VARCHAR(25)"`
	Phone     string    `json:"phone" gorm:"type: VARCHAR(255)"`
	Address   string    `json:"address" gorm:"type: VARCHAR(255)"`
	Subscribe bool      `json:"subscribe" gorm:"default: false"`
	Role      string    `json:"role" gorm:"type: VARCHAR(25)"`
	Thumbnail string    `json:"thumbnail" gorm:"type: VARCHAR(255)"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
