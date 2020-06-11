package app

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HomeController struct{}

func (HomeController) index(c echo.Context) error {
	return c.String(http.StatusOK, "pong!")
}
