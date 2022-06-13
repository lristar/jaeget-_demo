package demo

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func NewHTTPServe() (*gin.Engine, func() error, error) {
	app := gin.Default()
	return app, Closed, nil
}

func Closed() error {
	fmt.Println("httpServer closed")
	return nil
}
