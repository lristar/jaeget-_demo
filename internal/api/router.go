package api

import (
	"github.com/gin-gonic/gin"
)

var srv IHTTPSERVICE

func NewRouter(app *gin.Engine) *gin.Engine {
	root := app.Group("/")
	{
		root.GET("", srv.Hello)
		root.POST("", srv.Hay)
	}
	return app
}

// 中间件
