package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExampleHandler struct{}

func NewExampleHandler() *ExampleHandler {
	return &ExampleHandler{}
}

func (h *ExampleHandler) HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
