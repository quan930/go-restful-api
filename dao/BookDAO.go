package dao

import "bookapi/entity"

type BookDAO interface {
	SelectBooksAll() *[]entity.Book
	SelectBookById(id uint) *entity.Book
	InsertBook(book entity.Book) *entity.Book
	UpdateBookById(id uint, book entity.Book) *entity.Book
	DeleteBookById(id uint) *int64
}
