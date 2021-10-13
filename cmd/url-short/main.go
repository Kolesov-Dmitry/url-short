package main

import (
	"log"
	"url-short/internal/app"
)

// go build -o bin\url-short.exe cmd\url-short\main.go
func main() {
	a := app.NewApplication()
	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
