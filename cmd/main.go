package main

import (
	"log"

	"github.com/stazoloto/todo/internal/app"
)

func main() {
	app := app.NewApp()
	if err := app.Run(); err != nil {
		log.Fatalf("Failed to run application: %v", err)
	}
}
