package api

import (
	"ending/test/jaeget-_demo/internal/pkg"
	"github.com/gin-gonic/gin"
)

var Srv IHTTPSERVICE

func NewRouter(app *gin.Engine) *gin.Engine {
	root := app.Group("/")
	{
		root.GET("", pkg.HandleGin(Srv.Hello))
		root.POST("", pkg.HandleGin(Srv.Hay))
	}
	return app
}

// 中间件
