package main

import (
	"bookapi/config"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	config.RouteConfig(router)
	router.Run("localhost:8080")
}