package controller

import (
	"bookapi/entity"
	"bookapi/service"
	"fmt"
	"github.com/emicklei/go-restful"
	"net/http"
)

type BookCon struct {
}

var bookService service.BookService

func init() {
	fmt.Println("bookService init")
	bookService = new(service.BookServiceImpl)
}

func (receiver BookCon) GetBooks(request *restful.Request, response *restful.Response) {
	response.WriteEntity(bookService.GetList())
}

func (receiver BookCon) GetBookByID(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("id")
	book := bookService.GetBookById(id)
	if book == nil {
		response.WriteErrorString(http.StatusNotFound, "book not found")
	} else {
		response.WriteEntity(book)
	}
}

func (receiver BookCon) AddBook(request *restful.Request, response *restful.Response) {
	bookAO := new(entity.BookAO)
	err := request.ReadEntity(&bookAO)
	if err == nil {
		book := bookService.AddBook(*bookAO)
		if book == nil {
			response.WriteErrorString(http.StatusInternalServerError, "参数异常")
		} else {
			response.WriteHeaderAndEntity(http.StatusCreated, book)
		}
	} else {
		response.WriteErrorString(http.StatusInternalServerError, "JSON异常")
	}
}

func (receiver BookCon) UpdateBook(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("id")
	var book entity.Book
	err := request.ReadEntity(&book)
	if err == nil {
		bookNew := bookService.UpdateBook(id, book)
		if bookNew == nil {
			response.WriteErrorString(http.StatusInternalServerError, "参数异常")
		} else {
			response.WriteHeaderAndEntity(http.StatusCreated, bookNew)
		}
	} else {
		response.WriteErrorString(http.StatusInternalServerError, "JSON异常")
	}
}

func (receiver BookCon) DeleteBook(request *restful.Request, response *restful.Response) {
	//todo 删除完善
	id := request.PathParameter("id")
	if len(id) == 0 {
		response.WriteErrorString(http.StatusInternalServerError, "参数异常")
		return
	}
	book := bookService.DeleteBook(id)
	if book == nil {
		response.WriteErrorString(http.StatusInternalServerError, "book not found,can't delete")
	} else {
		response.WriteEntity(*book)
	}
}
