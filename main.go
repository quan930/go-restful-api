// testProject API.
//
// 这里是一些简单的描述
//
// Terms Of Service:
//
// 描述使用接口服务的一些协议
//
//     Schemes: http, https
//     Host: localhost
//     BasePath: /v1
//     Version: 0.0.1
//     Contact: lilq<lilq@test.com>
//
//     Consumes:
//     - application/json
//     - application/xml
//
//     Produces:
//     - application/json
//     - application/xml
//
// swagger:meta
package main

import (
	"bookapi/config"
	"github.com/emicklei/go-restful"
	"log"
	"net/http"
)

func main() {
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})
	//Register
	config.Register(wsContainer)
	log.Printf("start listening on localhost:8080")
	server := &http.Server{Addr: ":8080", Handler: wsContainer}
	log.Fatal(server.ListenAndServe())
}
