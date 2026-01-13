package main

import (
	"go_2601_04/internal/application/message"
	"go_2601_04/internal/interface/http"

	"github.com/gin-gonic/gin"
)

func main() {
	appSvc := message.NewService()

	handler := http.NewHandler(appSvc)

	r := gin.Default()
	r.GET("/hello", handler.HelloWorld)

	r.Run(":8080")
}
