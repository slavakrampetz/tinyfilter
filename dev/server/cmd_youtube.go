package server

import (
	"errors"
	"github.com/labstack/echo"
	"net/http"
	"tinyfilter/dev/command/youtube"
	"tinyfilter/dev/log"
)

// Youtube
// noinspection GoUnusedParameter
func cmdYoutube(c echo.Context) error {

	state := c.Param("state")

	isOn := true
	switch state {
	case "get":
		return cmdYoutubeRead(c)
	case "off", "no":
		isOn = false
	case "on", "yes":
		isOn = true
	default:
		return errors.New("unknown state parameter: " + state)
	}

	log.Inf("Youtube command…")
	err := youtube.Exec(isOn)
	if err != nil {
		log.Err(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	log.Inf("success: Youtube status set to", isOn)

	return c.String(http.StatusOK, "Youtube: OK")
}

func cmdYoutubeRead(c echo.Context) error {
	status, err := youtube.ExecRead()
	if err != nil {
		log.Err(err)
		return c.NoContent(http.StatusInternalServerError)
	}
	return c.JSON(http.StatusOK, uint8(status))
}
