package main

import (
	"log"
	"os"
	"web-monitor/app"
)

// go run main.go ip --host amazon.co.uk
func main() {
	app := app.Gerar()
	if error := app.Run(os.Args); error != nil {
		log.Fatal(error)
	}
}
