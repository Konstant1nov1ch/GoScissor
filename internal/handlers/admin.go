package handlers

import (
	"GoScissor/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

func Admin(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokens []models.Token
		if err := db.Find(&tokens).Error; err != nil {
			c.String(http.StatusInternalServerError, "500 internal server error")
			return
		}

		c.HTML(http.StatusOK, "admin.html", gin.H{
			"title":  "Admin",
			"tokens": tokens,
		})
	}
}
