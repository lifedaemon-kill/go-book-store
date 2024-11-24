package pkg

import "go-book-store/models"

type BookRepository interface {
	Create(book models.Book) (int, error)
	Update(bookId int64, book models.Book) error
	Delete(id int64) error

	GetAll(skip, limit int) ([]models.Book, error)
	GetById(id int64) (models.Book, error)
}
