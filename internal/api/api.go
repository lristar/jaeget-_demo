package api

import (
	"ending/test/jaeget-_demo/internal/pkg"
)

type IHTTPSERVICE interface {
	Hello(ctx *pkg.Context)
	Hay(ctx *pkg.Context)
}
