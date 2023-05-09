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

func (t *Token) BeforeCreate(scope *gorm.Scope) error {
	for {
		shortURL := generateShortURL()
		var token Token
		if result := scope.DB().Where(&Token{ShortURL: shortURL}).First(&token); result.Error != nil {
			if result.RecordNotFound() {
				t.ShortURL = shortURL
				break
			} else {
				return result.Error
			}
		}
	}
	return nil
}

func generateShortURL() string {
	characters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	length := 7
	b := make([]rune, length)
	for i := range b {
		b[i] = characters[rand.Intn(len(characters))]
	}
	return string(b)
}
