package server

import (
	"github.com/labstack/echo"
)

// Ping
// noinspection GoUnusedParameter
func ping(c echo.Context) error {
	return c.String(200, "pong")
}

// Home, not a thing
func home(c echo.Context) error {
	return c.String(401, "You are not welcome here")
}
