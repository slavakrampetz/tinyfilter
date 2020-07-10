package app

import (
	"context"
	"errors"
	"fmt"
	e "github.com/labstack/echo"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"tinyfilter/dev/etc"
	"tinyfilter/dev/server"
)

// Application structure
type Application struct {

	// Port for listen
	Port int

	// HTTP server instance
	Server *e.Echo
	// Logger shortcut
	Log e.Logger
}

// Get address of server
func (app *Application) GetAddress() string {
	return ":" + strconv.Itoa(app.Port)
}


// Init HTTP server
func (app *Application) Init() {
	app.Server = server.CreateEcho()
	app.Log = app.Server.Logger
}

func (app *Application) ReloadConfig() {
	fmt.Println("reload config...")
}

// Launch and listen to OS signals
func (app *Application) Start() (chan bool, error) {

	app.Init()

	isOk := true
	go func() {
		if err := app.Server.Start(app.GetAddress()); err != nil {
			isOk = false
			app.Log.Fatal("shutting down...")
		}
	}()

	if !isOk {
		return nil, errors.New("cannot start HTTP server")
	}

	done := make(chan bool, 1)

	// OS Signals
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGKILL)

	go func() {
		for sig := range signals {
			switch sig {
			case syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL:
				fmt.Println(sig, " signal received, exiting...")
				done <- true
				return
			case syscall.SIGHUP:
				app.ReloadConfig()
			}
		}
	}()

	return done, nil
}

// Stop
func (app *Application) Stop() {

	// Shutdown context
	ctx, cancel := context.WithTimeout(context.Background(), etc.DefaultTimeoutShutdown)
	defer cancel()

	if err := app.Server.Shutdown(ctx); err != nil {
		app.Log.Fatal(err)
	}
}

