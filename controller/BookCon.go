package controller

import (
	"bookapi/entity"
	"bookapi/service"
	"bookapi/util"
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/go-playground/validator/v10"
	"log"
)

type BookCon struct {
}

var bookService service.BookService
//参数校验
var validate *validator.Validate

func init() {
	fmt.Println("bookService init")
	bookService = new(service.BookServiceImpl)
	validate = validator.New()
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
// Responses:
//  200: response
func (receiver BookCon) GetBooks(request *restful.Request, response *restful.Response) {
	books := bookService.GetList()
	responseBody := new(entity.Response)
	if books==nil {
		responseBody.Body.Code = 400
		responseBody.Body.Msg = "业务异常"
		responseBody.Body.Data = nil
		response.WriteEntity(responseBody.Body)
	}else {
		responseBody.Body.Code = 200
		responseBody.Body.Msg = "successful"
		responseBody.Body.Data = books
		response.WriteEntity(responseBody.Body)
	}
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
// Responses:
//  200: response
//
func (receiver BookCon) GetBookByID(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("id")
	book := bookService.GetBookById(id)
	responseBody := new(entity.Response)
	if book == nil {
		responseBody.Body.Code = 404
		responseBody.Body.Msg = "not found"
		responseBody.Body.Data = nil
		response.WriteEntity(responseBody.Body)
	} else {
		responseBody.Body.Code = 200
		responseBody.Body.Msg = "successful"
		responseBody.Body.Data = book
		response.WriteEntity(responseBody.Body)
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
// Responses:
//  200: response
//
func (receiver BookCon) AddBook(request *restful.Request, response *restful.Response) {
	bookAO := new(entity.BookAO)
	err := request.ReadEntity(&bookAO)
	responseBody := new(entity.Response)
	if err == nil {
		// 参数校验
		err := validate.Struct(bookAO)
		if err != nil {
			responseBody.Body.Code = 500
			responseBody.Body.Msg = "参数异常:"+util.ValidateErrorFormat(err)
			response.WriteEntity(responseBody.Body)
			return
		}
		book := bookService.AddBook(*bookAO)
		if book == nil {
			responseBody.Body.Code = 404
			responseBody.Body.Msg = "参数异常"
			responseBody.Body.Data = nil
			response.WriteEntity(responseBody.Body)
		} else {
			responseBody.Body.Code = 201
			responseBody.Body.Msg = "successful"
			responseBody.Body.Data = book
			response.WriteEntity(responseBody.Body)
		}
	} else {
		responseBody.Body.Code = 500
		responseBody.Body.Msg = "JSON异常"
		responseBody.Body.Data = nil
		response.WriteEntity(responseBody.Body)
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
// Responses:
//  200: response
//
func (receiver BookCon) UpdateBook(request *restful.Request, response *restful.Response) {
	id := request.PathParameter("id")
	var book entity.BookUO
	err := request.ReadEntity(&book)
	log.Println(book)
	responseBody := new(entity.Response)
	if err == nil {
		// 参数校验
		err := validate.Struct(book)
		if err != nil {
			responseBody.Body.Code = 404
			responseBody.Body.Msg = "参数异常:"+util.ValidateErrorFormat(err)
			response.WriteEntity(responseBody.Body)
			return
		}
		bookNew := bookService.UpdateBook(id, book)
		if bookNew == nil {
			responseBody.Body.Code = 404
			responseBody.Body.Msg = "参数异常"
			responseBody.Body.Data = nil
			response.WriteEntity(responseBody.Body)
		} else {
			responseBody.Body.Code = 201
			responseBody.Body.Msg = "successful"
			responseBody.Body.Data = bookNew
			response.WriteEntity(responseBody.Body)
		}
	} else {
		responseBody.Body.Code = 500
		responseBody.Body.Msg = "JSON异常"
		responseBody.Body.Data = nil
		response.WriteEntity(responseBody.Body)
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
// Responses:
//  200: response
//
func (receiver BookCon) DeleteBook(request *restful.Request, response *restful.Response) {
	//todo 删除完善
	id := request.PathParameter("id")
	responseBody := new(entity.Response)
	if len(id) == 0 {
		responseBody.Body.Code = 500
		responseBody.Body.Msg = "参数异常"
		responseBody.Body.Data = nil
		response.WriteEntity(responseBody.Body)
		return
	}
	book := bookService.DeleteBook(id)
	if book == nil {
		responseBody.Body.Code = 500
		responseBody.Body.Msg = "book not found,can't delete"
		responseBody.Body.Data = nil
		response.WriteEntity(responseBody.Body)
	} else {
		responseBody.Body.Code = 200
		responseBody.Body.Msg = "successful"
		responseBody.Body.Data = book
		response.WriteEntity(responseBody.Body)
	}
}
