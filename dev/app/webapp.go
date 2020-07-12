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

func RunWeb() error {

	web := &WepApp{
		Port: etc.DefaultPort,
	}

	quit, err := web.Start()
	if err != nil {
		fmt.Println("Error starting web-server> ", err)
		return err
	}

	// Wait for interrupt signal to gracefully shutdown the server
	if quit != nil {
		<-quit
	}

	// Shutdown right now, with context timeout
	web.Stop()
	return nil
}

// WepApp structure
type WepApp struct {

	// Port for listen
	Port int

	// HTTP server instance
	Server *e.Echo
	// Logger shortcut
	Log e.Logger
}

// Get address of server
func (web *WepApp) getAddress() string {
	return ":" + strconv.Itoa(web.Port)
}

func (web *WepApp) ReloadConfig() {
	fmt.Println("TODO: reload config...")
}

// Launch and listen to OS signals
func (web *WepApp) Start() (chan bool, error) {

	web.Server = server.CreateEcho()
	web.Log = web.Server.Logger

	isOk := true
	go func() {
		if err := web.Server.Start(web.getAddress()); err != nil {
			isOk = false
			web.Log.Fatal("shutting down...")
		}
	}()

	if !isOk {
		return nil, errors.New("cannot start HTTP server")
	}

	done := make(chan bool, 1)

	// OS Signals
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	go func() {
		for sig := range signals {
			switch sig {
			case syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL:
				fmt.Println(sig, " signal received, exiting...")
				done <- true
				return
			case syscall.SIGHUP:
				web.ReloadConfig()
			}
		}
	}()

	return done, nil
}

// Stop
func (web *WepApp) Stop() {

	// Shutdown context
	ctx, cancel := context.WithTimeout(context.Background(), etc.DefaultTimeoutShutdown)
	defer cancel()

	if err := web.Server.Shutdown(ctx); err != nil {
		web.Log.Fatal(err)
	}
}
