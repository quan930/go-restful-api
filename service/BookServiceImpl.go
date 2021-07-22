package service

import (
	"bookapi/entity"
	"fmt"
	"github.com/satori/go.uuid"
)

type BookServiceImpl struct {

}

var books = []entity.Book{
	{ID: "1", Name: "Thinking in JAVA", Price: 56.99, Author: "Bruce Eckel"},
	{ID: "2", Name: "重构", Price: 17.99, Author: "马丁福勒"},
	{ID: "3", Name: "core JAVA", Price: 39.99, Author: "Gary Cornell"},
}

func (receiver BookServiceImpl) GetList() *[]entity.Book {
	//fmt.Println(books)
	return &books
}

func (receiver BookServiceImpl) GetBookById(id string) *entity.Book{
	for i := range books {
		if books[i].ID==id {
			return &books[i]
		}
	}
	return nil
}

func (receiver BookServiceImpl) AddBook(book entity.Book) *entity.Book {
	fmt.Print("增加")
	book.ToString()
	book.ID = uuid.Must(uuid.NewV4()).String()
	books = append(books, book)
	return &book
}

func (receiver BookServiceImpl) UpdateBook(id string,book entity.Book) *entity.Book {
	fmt.Print("更新:"+id+"\t")
	book.ToString()
	//update
	for i := range books {
		if books[i].ID==id {
			if len(book.Name)!=0 {
				books[i].Name = book.Name
			}
			if len(book.Author)!=0 {
				books[i].Author = book.Author
			}
			if book.Price!=0 {
				books[i].Price = book.Price
			}
			return &books[i]
		}
	}
	return nil
}

func (receiver BookServiceImpl) DeleteBook(id string) *entity.Book {
	fmt.Println("删除:"+string(id))
	for i := range books {
		if books[i].ID == id {
			src := books[i]
			books = append(books[:i], books[i+1:]...)
			return &src
		}
	}
	return nil
}