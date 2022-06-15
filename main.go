package main

import (
	"jaegerDemo/app"
)

func main() {
	svc, closes, err := app.New()
	if err != nil {
		closes.CloseAll()
		return
	}
	svc.Run()
}
