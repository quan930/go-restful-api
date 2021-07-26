package service

import (
	"bookapi/dao"
	"bookapi/entity"
	"fmt"
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
		fmt.Println(err)
	}
	return bookDAO.SelectBookById(uint(idUit))
}

func (receiver BookServiceImpl) AddBook(book entity.BookAO) *entity.Book {
	fmt.Print("增加")
	bookNew := new(entity.Book)
	bookNew.Name = book.Name
	bookNew.Price = book.Price
	bookNew.Author = book.Author
	bookNew.ToString()
	return bookDAO.InsertBook(*bookNew)
}

func (receiver BookServiceImpl) UpdateBook(id string, book entity.Book) *entity.Book {
	fmt.Print("更新:" + id + "\t")
	book.ToString()
	idUit, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return bookDAO.UpdateBookById(uint(idUit), book)
}

func (receiver BookServiceImpl) DeleteBook(id string) *int64 {
	fmt.Println("删除:" + string(id))
	idUit, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		fmt.Println(err)
	}
	return bookDAO.DeleteBookById(uint(idUit))
}
