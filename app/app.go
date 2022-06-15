package app

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"jaegerDemo/internal/api"
	server "jaegerDemo/internal/server"
	service "jaegerDemo/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	"github.com/opentracing/opentracing-go"
)

type App struct {
	jae opentracing.Tracer
	g   *gin.Engine
}
type CloseMap []func() error

func (c CloseMap) CloseAll() {
	for _, v := range c {
		_ = v()
	}
	return
}

func New() (*App, CloseMap, error) {
	closes := CloseMap{}
	_, close1, err := service.NewService()
	if err != nil {
		return nil, closes, err
	}
	closes = append(closes, close1)
	jaeger, close2, err := server.NewJaeger("jaeger_demo")
	if err != nil {
		closes.CloseAll()
		return nil, closes, err
	}
	closes = append(closes, close2.Close)
	g, close3, err := server.NewHTTPServe()
	if err != nil {
		closes.CloseAll()
		return nil, closes, err
	}
	closes = append(closes, close3)
	return &App{
		jae: jaeger,
		g:   g,
	}, closes, nil
}

func (a *App) Run() {
	a.g.Use(ginhttp.Middleware(a.jae, ginhttp.OperationNameFunc(func(r *http.Request) string {
		return fmt.Sprintf("HTTP %s %s", r.Method, r.URL.Path)
	}), ginhttp.MWSpanObserver(func(span opentracing.Span, r *http.Request) {
		if r.URL.RawQuery != "" {
			span.SetTag("http.params", r.URL.RawQuery)
		}
		span.SetTag("http.remoteAddr", r.RemoteAddr)
		bts, _ := ioutil.ReadAll(r.Body)
		span.SetTag("http.body", string(bts))
		r.Body = ioutil.NopCloser(bytes.NewBuffer(bts))
	})))
	a.g = api.NewRouter(a.g)
	err := a.g.Run(":8080")
	if err != nil {
		panic(err)
	}
}
