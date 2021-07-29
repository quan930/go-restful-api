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

// swagger:route GET /books books getbooks
//
// 获取全部书籍
//
// 获取全部书籍(描述)
//     Consumes:
//     - application/json
//     Produces:
//     - application/json
//     Schemes: http
func (receiver BookCon) GetBooks(request *restful.Request, response *restful.Response) {
	response.WriteEntity(bookService.GetList())
}

// swagger:route GET /books/{Id} books getBookById
//
// 获取指定书籍
//
// 根据id获取指定书籍
// Consumes:
// - application/json
// Produces:
// - application/json
// Schemes: http
//
func (receiver BookCon) GetBookByID(request *restful.Request, response *restful.Response) {

	id := request.PathParameter("id")
	book := bookService.GetBookById(id)
	if book == nil {
		response.WriteErrorString(http.StatusNotFound, "book not found")
	} else {
		response.WriteEntity(book)
	}
}

// swagger:route POST /books books addBook
//
// 增加书籍
//
// 增加书籍(name,price,author)
// Consumes:
// - application/json
// Produces:
// - application/json
// Schemes: http
//
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

// swagger:route PATCH /books/{Id} books updateBook
//
// 修改书籍
//
// 增加书籍(name,price,author)
// Consumes:
// - application/json
// Produces:
// - application/json
// Schemes: http
//
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


// swagger:route DELETE /books/{Id} books deleteBookById
//
// 删除指定书籍
//
// 根据id删除指定书籍
// Consumes:
// - application/json
// Produces:
// - application/json
// Schemes: http
//
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
