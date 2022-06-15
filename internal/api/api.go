package api

import (
	"jaegerDemo/internal/pkg"
)

type IHTTPSERVICE interface {
	Hello(ctx *pkg.Context)
	Hay(ctx *pkg.Context)
}
