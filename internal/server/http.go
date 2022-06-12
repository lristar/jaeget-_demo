package demo

import (
	"github.com/gin-gonic/gin"
)

func NewHTTPServe() *gin.Engine {
	app := gin.Default()
	return app
}
