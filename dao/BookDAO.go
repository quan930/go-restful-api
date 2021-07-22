package dao

import "bookapi/entity"

type BookDAO interface {
	SelectBooksAll() *[]entity.Book
	SelectBookById(id string) *entity.Book
	InsertBook(book entity.Book) *entity.Book
	UpdateBookById(id string,book entity.Book) *entity.Book
	DeleteBookById(id string) *int64
}
