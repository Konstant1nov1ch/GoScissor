package main

import (
	. "GoScissor/internal/handlers"
	. "GoScissor/internal/models"
	"GoScissor/internal/pkg/cache"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"time"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=url_shortener sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	db.AutoMigrate(&Token{})

	cache := cache.NewCache(10, 100, time.Hour*24)

	router := gin.Default()
	// указываем директорию с шаблонами HTML
	router.LoadHTMLGlob("templates/*.html")

	// определяем маршрут для /admin/tokens
	router.GET("/admin/tokens", Admin(db))
	router.GET("/:short_url", Redirect(db, cache))
	router.POST("/admin/tokens", CreateToken(db, cache))

	err1 := router.Run(":8080")
	if err1 != nil {
		return
	}
}
