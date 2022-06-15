package demo

import (
	"fmt"
	"jaegerDemo/internal/api"
	"jaegerDemo/internal/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
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
