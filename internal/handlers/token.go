package handlers

import (
	"GoScissor/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func CreateToken(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			FullURL  string `json:"full_url" binding:"required"`
			ShortURL string `json:"short_url" binding:"required"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.String(http.StatusBadRequest, "400 bad request")
			return
		}

		token := models.Token{
			FullURL:  input.FullURL,
			ShortURL: input.ShortURL,
		}

		if err := db.Create(&token).Error; err != nil {
			c.String(http.StatusInternalServerError, "500 internal server error")
			return
		}

		c.JSON(http.StatusOK, gin.H{"data": token})
	}
}
