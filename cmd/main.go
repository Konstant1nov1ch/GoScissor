package main

import (
	. "GoScissor/internal/handlers"
	"GoScissor/internal/models"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=url_shortener sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&models.Token{})

	router := gin.Default()

	router.GET("/:short_url", Redirect(db))
	router.GET("/admin/tokens", Admin(db))
	router.POST("/admin/tokens", CreateToken(db))

	router.Run(":8080")
}
