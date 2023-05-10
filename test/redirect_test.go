package test

import (
	"GoScissor/internal/handlers"
	"GoScissor/internal/models"
	"GoScissor/internal/pkg/cache"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRedirectByShortURLHandler(t *testing.T) {
	db, err := gorm.Open("sqlite3", ":memory:")
	cache := cache.NewCache(10, 100, time.Minute)
	if err != nil {
		t.Fatalf("failed to open database: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&models.Token{})

	token := models.Token{FullURL: "https://example.com", ShortURL: "abcd"}
	db.Create(&token)

	r := gin.Default()
	r.GET("/:short_url", handlers.Redirect(db, cache))

	req, err := http.NewRequest("GET", "/"+token.ShortURL, nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp := httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusMovedPermanently, resp.Code)
	assert.Equal(t, "https://example.com", resp.Header().Get("Location"))

	// Повторный запрос должен быть обработан из кэша
	req, err = http.NewRequest("GET", "/"+token.ShortURL, nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp = httptest.NewRecorder()
	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusMovedPermanently, resp.Code)
	assert.Equal(t, "https://example.com", resp.Header().Get("Location"))

	// Запрос с неверным коротким URL должен возвращать ошибку 404
	req, err = http.NewRequest("GET", "/non-existing", nil)
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}

	resp = httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusNotFound, resp.Code)
}
