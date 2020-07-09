package main

import (
	"context"

	"os"
	"os/signal"

	"strconv"
	"time"

	"tinyfilter/dev/server"
)

const Port = 8085
const TimeoutShutdown = 10*time.Second

func main() {

	echo := server.CreateEcho()

	// Start
	go func() {
		if err := echo.Start(":" + strconv.Itoa(Port)); err != nil {
			echo.Logger.Info("shutting down...")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	// the server with a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), TimeoutShutdown)
	defer cancel()

	if err := echo.Shutdown(ctx); err != nil {
		echo.Logger.Fatal(err)
	}
}
