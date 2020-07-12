package server

import (
	// "path/filepath"
	// "pkg/mod/github.com/jordan-wright/unindexed@v0.0.0-20181209214434-78fa79113c0f"
	// "time"

	// Web framework
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	rand "github.com/labstack/gommon/random"
	"os"
)

const LogFormatCombined = `${time_rfc3339} ${id} ${remote_ip} ` +
	`${method} "${uri}"  ${status}:${error} ${latency} ${latency_human}` +
	` ${bytes_in}/${bytes_out}  ${user_agent}`

// Interface
func CreateEcho() *echo.Echo {

	handler := initEcho()

	// Set up our root handlers
	handler.GET("/", home)

	// Ping
	handler.GET("/ping", ping)
	handler.GET("/ping/", ping)
	//
	// // Set up our API
	// // handler.Mount("/api/v1/", v1.CreateHandler())

	// // Set up static file serving
	// staticPath, _ := filepath.Abs("../../static/")
	// fs := http.FileServer(unindexed.Dir(staticPath))
	// handler.Handle("/*", fs)

	return handler
}

// Ping
// noinspection GoUnusedParameter
func ping(c echo.Context) error {
	return c.String(200, "pong")
}

// Home, not a thing
func home(c echo.Context) error {
	return c.String(401, "You are not welcome here")
}

// initWeb router
func initEcho() *echo.Echo {

	e := echo.New()
	e.HideBanner = true

	// Logging
	e.Use(mw.LoggerWithConfig(mw.LoggerConfig{
		Format:           LogFormatCombined + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00",
		Output:           os.Stdout,
	}))

	// Request ID
	e.Use(mw.RequestIDWithConfig(mw.RequestIDConfig{
		Skipper:   nil,
		Generator: requestId,
	}))

	// Recover of panic errors
	e.Use(mw.Recover())

	// Not implemented?
	// handler.Use(mw.RealIP)
	// handler.Use(mw.Compress(???))
	// handler.Use(mw.Timeout(60 * time.Second))
	return e
}

//
// Helpers

func requestId() string {
	return rand.String(12)
}
