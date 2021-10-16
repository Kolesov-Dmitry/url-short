package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"
	"url-short/internal/app"
)

// go build -o bin\url-short.exe cmd\url-short\main.go
func main() {
	a := app.NewApplication()
	if err := a.Run(); err != nil {
		log.Fatal(err)
	}

	// Handle ctrl+c
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	<-ctx.Done()
	cancel()

	// Shutdown server
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	a.Close(ctx)
	cancel()
}
