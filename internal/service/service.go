package demo

import (
	"ending/test/jaeget-_demo/internal/api"
	"ending/test/jaeget-_demo/internal/pkg"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Service struct {
}

func NewService() (*Service, func() error, error) {
	s := &Service{}
	api.Srv = s
	return s, Closed, nil
}

func Closed() error {
	fmt.Println("SERVICE CLOSED")
	return nil
}

func (s *Service) Hello(ctx *pkg.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "hello"})
}

func (s *Service) Hay(ctx *pkg.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "hay"})
}

// 通过反射获取方法名字，请求和响应参数
