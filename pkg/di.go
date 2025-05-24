package pkg

import (
	"github.com/asm2212/go-gin-starter/internal/api"
	"github.com/asm2212/go-gin-starter/internal/service"
)

func InitializeHandler() *api.ExampleHandler {
	_ = service.NewExampleService() // In a real app, you'd pass this to the handler
	return api.NewExampleHandler()
}
