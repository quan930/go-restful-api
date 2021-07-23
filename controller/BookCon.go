package controller

import (
	"bookapi/entity"
	"bookapi/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var bookService service.BookService
func init() {
	fmt.Println("bookService init")
	bookService = new(service.BookServiceImpl)
}

func GetBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK,*bookService.GetList())
}

func GetBookByID(c *gin.Context) {
	id := c.Param("id")
	book := bookService.GetBookById(id)
	if book==nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "book not found"})
	}else {
		c.IndentedJSON(http.StatusOK, *book)
	}
}

func AddBook(c *gin.Context) {
	//todo 参数校验
	var book entity.BookAO

	if err := c.BindJSON(&book); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "JSON异常"})
		return
	}

	bookNew := bookService.AddBook(book)
	if bookNew==nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "参数异常"})
	}else {
		c.IndentedJSON(http.StatusCreated, *bookNew)
	}
}

func UpdateBook(c *gin.Context)  {
	//todo 参数校验
	id := c.Param("id")
	var book entity.Book
	if err := c.BindJSON(&book); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "JSON异常"})
		return
	}

	bookNew := bookService.UpdateBook(id, book)
	if bookNew==nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "参数异常"})
	}else {
		c.IndentedJSON(http.StatusCreated, *bookNew)
	}
}

func DeleteBook(c *gin.Context)  {
	id := c.Param("id")
	if len(id)==0{
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "参数异常"})
		return
	}

	book := bookService.DeleteBook(id)
	if book==nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "book not found,can't delete"})
	}else {
		c.IndentedJSON(http.StatusOK, *book)
	}
}