package main

import (
	. "GoScissor/internal/handlers"
	. "GoScissor/internal/models"
	"GoScissor/internal/pkg/cache"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
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
			log.Println(err)
		}
	}(db)

	db.AutoMigrate(&Token{})

	cache, err := cache.NewCache(10, 100, time.Minute*30)
	if err != nil {
		log.Println("failed to create cache: " + err.Error())
	}
	router := gin.Default()
	// указываем директорию с шаблонами HTML
	router.LoadHTMLGlob("web/templates/*.html")

	router.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "http://localhost:63342")
		c.Header("Access-Control-Allow-Methods", "POST, GET")
		c.Header("Access-Control-Allow-Headers", "Content-Type")
	})

	router.OPTIONS("/sci", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})
	//ToDO подумать о многопотоке

	// Добавляем обработчик для просмотра списка всех токенов из бд
	router.GET("/admin/tokens", Admin(db))
	// Добавляем обработчик для редиректа
	router.GET("/:short_url", Redirect(db, cache))
	// Добавляем обработчик для запросов POST на путь /sci
	router.POST("/sci", CreateToken(db, cache))

	err = router.Run(":8080")
	if err != nil {
		log.Println("Failed to create router : ", err)
	}
}
