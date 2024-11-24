package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files" // Это "swaggerFiles"
	"github.com/swaggo/gin-swagger"
	"go-book-store/db"
	_ "go-book-store/docs"
	"go-book-store/handlers"
	"go-book-store/logger"
)

// @title Book API
// @version 1.0
// @description API для управления книгами.
// @host localhost:8080

func main() {
	err := logger.Init("logger/file.log")
	if err != nil {
		panic(err)
	}
	db.InitDB()

	// Create a router
	r := gin.Default()

	// Настройка CORS для разрешения запросов с любого источника
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost:8080"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},  // Разрешить определенные методы
		AllowHeaders:     []string{"Content-Type", "Authorization"}, // Разрешить определенные заголовки
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	//Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Routes
	r.POST("/books/", handlers.CreateBook)
	r.GET("/books/", handlers.GetBooks)
	r.GET("/books/:book_id", handlers.GetBookById)
	r.PUT("/books/:book_id", handlers.UpdateBook)
	r.DELETE("/books/:book_id", handlers.DeleteBook)

	//Start server
	err = r.Run(":8080")
	if err != nil {
		return
	}
}
