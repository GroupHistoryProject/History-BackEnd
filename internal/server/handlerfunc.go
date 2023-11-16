package server

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func MapHandler(pathsToUrls map[string]string, fallback echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		URL, ok := pathsToUrls[c.Request().URL.Path]
		if ok {
			return c.Redirect(http.StatusMovedPermanently, URL)
		} else {
			return fallback(c)
		}
	}
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, world!")
}
