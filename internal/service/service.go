package demo

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Service struct {
}

func (s *Service) Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "hello"})
}

func (s *Service) Hay(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "hay"})
}

// 通过反射获取方法名字，请求和响应参数
