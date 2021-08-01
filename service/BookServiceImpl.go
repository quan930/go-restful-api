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
	log.Print("增加")
	bookNew := new(entity.Book)
	bookNew.Name = book.Name
	bookNew.Price = book.Price
	bookNew.Author = book.Author
	bookNew.ToString()
	return bookDAO.InsertBook(*bookNew)
}

func (receiver BookServiceImpl) UpdateBook(id string, book entity.BookUO) *entity.Book {
	log.Print("更新:" + id + "\t")
	bookNew := new(entity.Book)
	bookNew.Name = book.Name
	bookNew.Price = book.Price
	bookNew.Author = book.Author
	bookNew.ToString()
	idUit, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Println(err)
	}
	return bookDAO.UpdateBookById(uint(idUit), *bookNew)
}

func (receiver BookServiceImpl) DeleteBook(id string) *int64 {
	log.Println("删除:" + string(id))
	idUit, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Println(err)
	}
	return bookDAO.DeleteBookById(uint(idUit))
}
