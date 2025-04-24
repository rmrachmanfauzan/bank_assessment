package model

import "time"

type Account struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	NoRekening string    `json:"no_rekening"`
	Saldo     float64   `json:"saldo" default:"0"`
	UserID    uint      `json:"user_id"` // foreign key
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}	

type Tabung struct {
	no_rekening string
	saldo float64
}

type Tarik struct{
	no_rekening string
	nominal float64
}

