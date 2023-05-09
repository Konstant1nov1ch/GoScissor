package handlers

import (
	"GoScissor/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func Redirect(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		shortURL := c.Param("short_url")

		var token models.Token
		if err := db.Where(&models.Token{ShortURL: shortURL, IsActive: true}).First(&token).Error; err != nil {
			if gorm.IsRecordNotFoundError(err) {
				c.String(http.StatusNotFound, "404 page not found")
			} else {
				c.String(http.StatusInternalServerError, "500 internal server error")
			}
			return
		}

		token.RequestsCount += 1
		if err := db.Save(&token).Error; err != nil {
			c.String(http.StatusInternalServerError, "500 internal server error")
			return
		}

		c.Redirect(http.StatusMovedPermanently, token.FullURL)
	}
}
