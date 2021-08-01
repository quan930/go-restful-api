package controller

import (
	"bookapi/entity"
	"bookapi/service"
	"bookapi/util"
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
	log.Println("bookService init")
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
	if books==nil {
		response.WriteEntity(entity.NewResponse(500,"业务异常",nil).Body)
	}else {
		response.WriteEntity(entity.NewResponse(200,"successful",books).Body)
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
	if book == nil {
		response.WriteEntity(entity.NewResponse(404,"找不到资源",nil).Body)
	} else {
		response.WriteEntity(entity.NewResponse(200,"successful",book).Body)
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
	//responseBody := new(entity.Response)
	if err == nil {
		// 参数校验
		err := validate.Struct(bookAO)
		if err != nil {
			response.WriteEntity(entity.NewResponse(500,"参数异常:"+util.ValidateErrorFormat(err),nil).Body)
			return
		}
		book := bookService.AddBook(*bookAO)
		if book == nil {
			response.WriteEntity(entity.NewResponse(500,"业务异常",nil).Body)
			return
		} else {
			response.WriteEntity(entity.NewResponse(201,"successful",book).Body)
			return
		}
	} else {
		response.WriteEntity(entity.NewResponse(500,"JSON异常",nil).Body)
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
	if err == nil {
		// 参数校验
		err := validate.Struct(book)
		if err != nil {
			response.WriteEntity(entity.NewResponse(500,"参数异常:"+util.ValidateErrorFormat(err),nil).Body)
			return
		}
		bookNew := bookService.UpdateBook(id, book)
		if bookNew == nil {
			response.WriteEntity(entity.NewResponse(500,"业务异常",nil).Body)
		} else {
			response.WriteEntity(entity.NewResponse(201,"successful",bookNew).Body)
		}
	} else {
		response.WriteEntity(entity.NewResponse(500,"JSON异常",nil).Body)
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
	if len(id) == 0 {
		response.WriteEntity(entity.NewResponse(500,"id不为空",nil).Body)
		return
	}
	book := bookService.DeleteBook(id)
	if book == nil {
		response.WriteEntity(entity.NewResponse(500,"book 不存在，不能删除",nil).Body)
	} else {
		response.WriteEntity(entity.NewResponse(201,"successful",book).Body)
	}
}

