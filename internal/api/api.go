package api

import "github.com/gin-gonic/gin"

type IHTTPSERVICE interface {
	Hello(ctx *gin.Context)
	Hay(ctx *gin.Context)
}
