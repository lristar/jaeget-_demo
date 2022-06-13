package main

import (
	"ending/test/jaeget-_demo/app"
)

func main() {
	svc, closes, err := app.New()
	if err != nil {
		closes.CloseAll()
		return
	}
	svc.Run()
}
