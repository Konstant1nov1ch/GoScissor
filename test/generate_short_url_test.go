package test

import (
	"GoScissor/internal/cache"
	"GoScissor/internal/handlers"
	"GoScissor/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/stretchr/testify/assert"
)

func TestGenerateShortURLHandler(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	cache := cache.NewCache(10, 100, time.Minute)
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&models.Token{})

	r := gin.Default()
	r.POST("/admin/tokens", handlers.CreateToken(db, cache))

	body := map[string]string{
		"full_url": "https://t.me/Algoru_bot",
	}

	reqBody, err := json.Marshal(body)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	req, err := http.NewRequest("POST", "/admin/tokens", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var response struct {
		Data models.Token `json:"data"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}

	assert.Equal(t, "https://t.me/Algoru_bot", response.Data.FullURL)
	assert.NotEmpty(t, response.Data.ShortURL)
}
