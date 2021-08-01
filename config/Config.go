package config

import (
	"bookapi/controller"
	"github.com/emicklei/go-restful"
	"log"
	"os"
	"strings"
	"time"
)

var logger = log.New(os.Stdout, "", 0)

func Register(container *restful.Container) {

	ws := new(restful.WebService)
	bookCon := new(controller.BookCon)
	cors := restful.CrossOriginResourceSharing{
		ExposeHeaders:  []string{"X-My-Header"},
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH"},
		CookiesAllowed: false,
		Container: container}

	//容器过滤器 跨域
	container.Filter(cors.Filter)
	//容器过滤器 跨域 配置OPTIONS 请求
	container.Filter(container.OPTIONSFilter)
	ws.
		Path("/api/v1/books").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_XML) // you can specify this per route as well

	// WebService过滤器 日志
	ws.Filter(NCSACommonLogFormatLogger())

	ws.Route(ws.GET("").To(bookCon.GetBooks))
	ws.Route(ws.GET("/{id}").To(bookCon.GetBookByID))
	ws.Route(ws.POST("").To(bookCon.AddBook))
	ws.Route(ws.PATCH("/{id}").To(bookCon.UpdateBook))
	ws.Route(ws.DELETE("/{id}").To(bookCon.DeleteBook))

	container.Add(ws)
}


func NCSACommonLogFormatLogger() restful.FilterFunction {
	return func(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
		var username = "-"
		if req.Request.URL.User != nil {
			if name := req.Request.URL.User.Username(); name != "" {
				username = name
			}
		}
		chain.ProcessFilter(req, resp)
		logger.Printf("%s - %s [%s] \"%s %s %s\" %d %d",
			strings.Split(req.Request.RemoteAddr, ":")[0],
			username,
			time.Now().Format("02/Jan/2006:15:04:05 -0700"),
			req.Request.Method,
			req.Request.URL.RequestURI(),
			req.Request.Proto,
			resp.StatusCode(),
			resp.ContentLength(),
		)
	}
}