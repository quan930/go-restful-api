package service

import (
	"bookapi/dao"
	"bookapi/entity"
	"fmt"
	"github.com/satori/go.uuid"
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

func (receiver BookServiceImpl) GetBookById(id string) *entity.Book{
	return bookDAO.SelectBookById(id)
}

func (receiver BookServiceImpl) AddBook(book entity.Book) *entity.Book {
	fmt.Print("增加")
	book.ToString()
	book.ID = uuid.Must(uuid.NewV4()).String()
	return bookDAO.InsertBook(book)
}

func (receiver BookServiceImpl) UpdateBook(id string,book entity.Book) *entity.Book {
	fmt.Print("更新:"+id+"\t")
	book.ToString()
	return bookDAO.UpdateBookById(id,book)
}

func (receiver BookServiceImpl) DeleteBook(id string) *int64 {
	fmt.Println("删除:"+string(id))
	return bookDAO.DeleteBookById(id)
}