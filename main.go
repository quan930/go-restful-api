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
