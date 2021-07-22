package config

import (
	"bookapi/controller"
	"github.com/gin-gonic/gin"
)

func RouteConfig(router *gin.Engine)  {
	router.GET("/books", controller.GetBooks)
	router.GET("/books/:id", controller.GetBookByID)
	router.POST("/books",controller.AddBook)
	router.PUT("/books/:id", controller.UpdateBook)
	router.DELETE("/books/:id",controller.DeleteBook)
}