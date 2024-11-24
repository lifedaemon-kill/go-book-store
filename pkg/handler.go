package pkg

import (
	"github.com/gin-gonic/gin"
	"go-book-store/db"
	"go-book-store/logger"
	"go-book-store/models"
	"net/http"
	"strconv"
)

// ErrorResponse структура для ошибок
type ErrorResponse struct {
	Error string `json:"error"`
}

// MessageResponse структура для сообщений
type MessageResponse struct {
	Message string `json:"message"`
}

// CreateBook
// @Summary Создать новую книгу
// @Description Добавляет новую книгу в базу данных.
// @Tags books
// @Accept json
// @Produce json
// @Param book body models.Book true "Данные книги"
// @Success 201 {object} models.Book
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /books [post]
func CreateBook(c *gin.Context) {
	var book models.Book

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	if err := db.DB.Create(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to create book"})
		return
	}
	logger.Log.Info("Created book ", book.Title)
	c.JSON(http.StatusCreated, book)
}

// GetBooks
// @Summary Получить список книг
// @Description Возвращает массив книг с возможностью пагинации.
// @Tags books
// @Accept json
// @Produce json
// @Param skip query int false "Смещение" default(0)
// @Param limit query int false "Лимит" default(10)
// @Success 200 {array} models.Book
// @Failure 500 {object} ErrorResponse
// @Router /books [get]
func GetBooks(c *gin.Context) {
	var books []models.Book
	skip, _ := strconv.Atoi(c.DefaultQuery("skip", "0"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	if err := db.DB.Offset(skip).Limit(limit).Find(&books).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to get books"})
		return
	}
	logger.Log.Info("Fetched books")
	c.JSON(http.StatusOK, books)
}

// GetBookById
// @Summary Получить информацию о книге
// @Description Возвращает информацию о книге по её ID.
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "ID книги"
// @Success 200 {object} models.Book
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /books/{id} [get]
func GetBookById(c *gin.Context) {
	var book models.Book
	id := c.Param("book_id")

	if _, err := strconv.Atoi(id); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: "Invalid ID format"})
		return
	}

	if err := db.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Book not found"})
		logger.Log.Warn("Attempted to fetch non-existent book with ID: ", id)
		return
	}
	logger.Log.Info("Fetched book ", book.Title)
	c.JSON(http.StatusOK, book)
}

// UpdateBook
// @Summary Обновить информацию о книге
// @Description Обновляет информацию о книге в базе данных.
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "ID книги"
// @Param book body models.Book true "Данные книги"
// @Success 200 {object} models.Book
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /books/{id} [put]
func UpdateBook(c *gin.Context) {
	var book models.Book
	id := c.Param("book_id")

	if err := db.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Book not found"})
		return
	}

	if err := c.ShouldBindJSON(&book); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	if err := db.DB.Save(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to update book"})
		return
	}
	logger.Log.Info("Updated book ", book.Title)
	c.JSON(http.StatusOK, book)
}

// DeleteBook
// @Summary Удалить книгу
// @Description Удаляет книгу из базы данных по её ID.
// @Tags books
// @Accept json
// @Produce json
// @Param id path int true "ID книги"
// @Success 200 {object} MessageResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /books/{id} [delete]
func DeleteBook(c *gin.Context) {
	var book models.Book
	id := c.Param("book_id")

	if err := db.DB.First(&book, id).Error; err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "Book not found"})
		return
	}

	if err := db.DB.Delete(&book).Error; err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: "Failed to delete book"})
		return
	}
	logger.Log.Info("Deleted book ", book.Title)
	c.JSON(http.StatusOK, MessageResponse{Message: "Book deleted successfully"})
}
