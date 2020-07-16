package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"

	"tinyfilter/dev/command/reload"
)

// Ping
// noinspection GoUnusedParameter
func cmdReload(c echo.Context) error {

	log.Info("Reloading…")
	err := reload.Exec()
	if err != nil {
		log.Error(err)
		return c.NoContent(http.StatusInternalServerError)
	}

	return c.String(http.StatusOK, "Reload: OK")
}
