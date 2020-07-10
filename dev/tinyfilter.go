package main

import (
	"fmt"
	"os"
	A "tinyfilter/dev/app"
	"tinyfilter/dev/etc"
)


func main() {

	app := &A.Application{
		Port: etc.DefaultPort,
	}

	quit, err := app.Start()
	if err != nil {
		fmt.Println("Error starting", err)
		os.Exit(2)
	}

	// Wait for interrupt signal to gracefully shutdown the server
	<-quit

	// Shutdown right now, with context timeout
	app.Stop()
}
