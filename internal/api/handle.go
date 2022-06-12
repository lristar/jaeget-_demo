package api

import (
	"jaegerDemo/internal/pkg"

	"github.com/gin-gonic/gin"
)

func HandleGin(f func(a *pkg.Context)) func(*gin.Context) {

	return func(g *gin.Context) {

	}
}
