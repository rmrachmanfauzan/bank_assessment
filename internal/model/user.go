package model

import (
	"time"
)

type User struct {
	ID		uint      `gorm:"primaryKey" json:"id"`
	Name  string `json:"name" validate:"required,min=3,max=100"`
	NIK   string `json:"nik" gorm:"not null" validate:"required,min=16,max=16"`
	Phone string `json:"phone" gorm:"not null" validate:"required,min=10,max=13"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Accounts  []Account `gorm:"foreignKey:UserID" json:"accounts"`
}

	





