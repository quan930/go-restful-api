package service

import "bookapi/entity"

type BookService interface {
	GetList() *[]entity.Book
	GetBookById(id string) *entity.Book
	AddBook(book entity.Book) *entity.Book
	UpdateBook(id string,book entity.Book) *entity.Book
	DeleteBook(id string) *entity.Book
}