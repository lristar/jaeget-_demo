package pkg

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
)

type Context struct {
	*gin.Context
}

func HandleGin(f func(a *Context)) func(*gin.Context) {
	return func(g *gin.Context) {
		c := New()
		c.Context = g
		f(c)
	}
}

func New() *Context {
	c := new(Context)
	gctx := new(gin.Context)
	gctx.Request = new(http.Request)
	gctx.Request.URL = new(url.URL)
	gctx.Params = make(gin.Params, 0)
	return c
}
