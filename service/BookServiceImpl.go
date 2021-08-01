package service

import (
	"bookapi/dao"
	"bookapi/entity"
	"log"
	"strconv"
)

type BookServiceImpl struct {
}

var bookDAO dao.BookDAO

func init() {
	bookDAO = new(dao.BookDAOImpl)
}

func (receiver BookServiceImpl) GetList() *[]entity.Book {
	return bookDAO.SelectBooksAll()
}

func (receiver BookServiceImpl) GetBookById(id string) *entity.Book {
	idUit, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Println(err)
	}
	return bookDAO.SelectBookById(uint(idUit))
}

func (receiver BookServiceImpl) AddBook(book entity.BookAO) *entity.Book {
	log.Print("book增加----book:")
	log.Print(book)
	bookNew := new(entity.Book)
	bookNew.Name = book.Name
	bookNew.Price = book.Price
	bookNew.Author = book.Author
	log.Println("\tbookNew:" + bookNew.ToString())
	return bookDAO.InsertBook(*bookNew)
}

func (receiver BookServiceImpl) UpdateBook(id string, book entity.BookUO) *entity.Book {
	log.Print("book更新---id:" + id + "\tbook:")
	log.Print(book)
	bookNew := new(entity.Book)
	bookNew.Name = book.Name
	bookNew.Price = book.Price
	bookNew.Author = book.Author
	log.Println("\tbookNew:"+bookNew.ToString())
	idUit, _ := strconv.ParseUint(id, 10, 64)
	return bookDAO.UpdateBookById(uint(idUit), *bookNew)
}

func (receiver BookServiceImpl) DeleteBook(id string) *int64 {
	log.Println("book删除---id:" + string(id))
	idUit, _ := strconv.ParseUint(id, 10, 64)
	return bookDAO.DeleteBookById(uint(idUit))
}
