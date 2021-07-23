package config

import (
	"bookapi/controller"
	"github.com/emicklei/go-restful"
)

func Register(container *restful.Container) {
	ws := new(restful.WebService)
	bookCon := new(controller.BookCon)
	ws.
		Path("/books").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

	ws.Route(ws.GET("").To(bookCon.GetBooks))
	ws.Route(ws.GET("/{id}").To(bookCon.GetBookByID))
	ws.Route(ws.POST("").To(bookCon.AddBook))
	ws.Route(ws.PATCH("/{id}").To(bookCon.UpdateBook))
	ws.Route(ws.DELETE("/{id}").To(bookCon.DeleteBook))

	container.Add(ws)
}
