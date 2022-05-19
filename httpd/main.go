package main

import (
	"ApiRest/httpd/handlers"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()
	server.GET("/", handlers.RequestHandler)
	err := server.Run()
	if err != nil {
		return
	}

}
