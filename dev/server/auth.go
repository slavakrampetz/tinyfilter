package server

import (
	"errors"
	"github.com/labstack/echo"
	mw "github.com/labstack/echo/middleware"
	"tinyfilter/dev/etc"

	"github.com/labstack/gommon/log"
)

var middlewareAuth echo.MiddlewareFunc

func getAuth() echo.MiddlewareFunc {

	if middlewareAuth != nil {
		return middlewareAuth
	}

	err := etc.Config.Read()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// head | query
	lookup := mw.DefaultKeyAuthConfig.KeyLookup
	if etc.Config.Auth.Type == "query" {
		lookup = "query:key"
	}

	middlewareAuth = mw.KeyAuthWithConfig(mw.KeyAuthConfig{
		KeyLookup: lookup,
		Validator: keyAuthenticate,
	})

	return middlewareAuth
}

// mw.KeyAuth(func(key string, c echo.Context) (bool, error) {
// 	log.Debug("Key received", key)
// 	return true, nil
// })

func keyAuthenticate(key string, _ echo.Context) (bool, error) {
	if etc.Config.Auth.Key != key {
		return false, errors.New("incorrect key specified")
	}
	log.Info("Auth key recognized:", key)
	return true, nil
}
