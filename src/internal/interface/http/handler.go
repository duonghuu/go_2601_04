package http

import (
	"go_2601_04/internal/application/message"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	appService message.Service
}

func NewHandler(s message.Service) Handler {
	return Handler{appService: s}
}

func (h Handler) HelloWorld(c *gin.Context) {
	name := c.DefaultQuery("name", "World")
	result := h.appService.GetHelloMessage(name)

	c.JSON(http.StatusOK, gin.H{
		"message": result,
	})
}
