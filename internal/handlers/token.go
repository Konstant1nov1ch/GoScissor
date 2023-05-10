package handlers

import (
	"GoScissor/internal/models"
	"GoScissor/internal/pkg/cache"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func CreateToken(db *gorm.DB, cache *cache.Cache) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			FullURL string `json:"full_url" binding:"required"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.String(http.StatusBadRequest, "400 bad request")
			return
		}

		token := models.Token{
			FullURL:  input.FullURL,
			ShortURL: models.GenerateShortURL(),
		}

		// Сохраните токен в базе данных
		if err := db.Create(&token).Error; err != nil {
			c.String(http.StatusInternalServerError, "500 internal server error")
			return
		}

		// Сохраните сокращенную ссылку в кэше
		cache.Set(input.FullURL, token.ShortURL)

		c.JSON(http.StatusOK, gin.H{"data": token})
	}
}
