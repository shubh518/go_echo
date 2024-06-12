package handlers

import (
	"net/http"

	"github.com/labstack/echo"
)

func MoviesHandler(c echo.Context) error {
	return c.String(http.StatusOK, "Ok!")
}
