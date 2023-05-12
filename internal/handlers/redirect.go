package handlers

import (
	"GoScissor/internal/models"
	"GoScissor/internal/pkg/cache"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func Redirect(db *gorm.DB, cache *cache.Cache) gin.HandlerFunc {
	return func(c *gin.Context) {
		shortURL := c.Param("short_url")

		// Пытаемся получить full_url по short_url из кэша
		if fullURL := cache.Get(shortURL); fullURL != nil {
			fmt.Println("Hello from 2Q Cache!")
			c.Redirect(http.StatusTemporaryRedirect, fullURL.(string))
			return
		}

		var token models.Token
		if err := db.Where(&models.Token{ShortURL: shortURL, IsActive: true}).First(&token).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				c.JSON(http.StatusNotFound, gin.H{"error": "short URL not found"})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			}
			return
		}

		token.RequestsCount += 1
		if err := db.Save(&token).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		// Обновляем кэш при редиректе
		cache.Set(shortURL, token.FullURL)

		c.Redirect(http.StatusTemporaryRedirect, token.FullURL)
	}
}
