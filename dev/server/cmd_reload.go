package server

import (
	"github.com/labstack/echo"
	"net/http"
	"tinyfilter/dev/log"

	"tinyfilter/dev/command/reload"
)

// Ping
// noinspection GoUnusedParameter
func cmdReload(c echo.Context) error {

	log.Inf("Reloading…")
	err := reload.Exec()
	if err != nil {
		log.Err(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.String(http.StatusOK, "Reload: OK")
}
