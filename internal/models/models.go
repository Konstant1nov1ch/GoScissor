package models

import (
	"github.com/jinzhu/gorm"
	"math/rand"
)

type Token struct {
	gorm.Model
	FullURL       string `gorm:"unique;not null"`
	ShortURL      string `gorm:"unique;not null;index"`
	RequestsCount int    `gorm:"default:0"`
	IsActive      bool   `gorm:"default:true"`
}

func GenerateShortURL() string {
	characters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	length := 7
	b := make([]rune, length)
	for i := range b {
		b[i] = characters[rand.Intn(len(characters))]
	}
	return string(b)
}
