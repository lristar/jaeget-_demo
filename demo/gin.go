package demo

import (
	"github.com/gin-gonic/gin"
)

func NewHTTPServe() {
	app := gin.Default()
	NewRouter(app)

}

func NewRouter(app *gin.Engine) {
	root := app.Group("/")
	{
		root.GET("", hello)
		root.POST("", hay)
	}
}
