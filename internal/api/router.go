package api

import (
	"jaegerDemo/internal/pkg"

	"github.com/gin-gonic/gin"
)

var Srv IHTTPSERVICE

func NewRouter(app *gin.Engine) *gin.Engine {
	root := app.Group("/")
	{
		root.GET("hello", pkg.HandleGin(Srv.Hello))
		root.GET("hay", pkg.HandleGin(Srv.Hay))
	}
	return app
}

// 中间件
