package main

import (
	"log"
	"url-short/internal/app"
)

func main() {
	a := app.NewApplication()
	if err := a.Run(); err != nil {
		log.Fatal(err)
	}
}
