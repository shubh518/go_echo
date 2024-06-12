package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"main.go/cmd/api/service"
)

// in this we handle the error
func PostIndexHandler(c echo.Context) error {
	data, err := service.GetAll()

	if err != nil {
		c.String(http.StatusBadGateway, "Unable to  process data")

	}
	// here we create a  new variable called res and we use to make keyword to create a map string any
	// and we will save res status ok
	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return c.JSON(http.StatusOK, res)

}

func PostSingleHandler(c echo.Context) error {
	id := c.Param("id")

	// Atoi is a function that converts a string to an integer
	// It takes a string as an argument and returns an integer.
	// It returns an error if the string is not a valid representation of an integer.

	idx, err := strconv.Atoi(id)

	if err != nil {
		c.String(http.StatusBadGateway, "Unable to  process data")

	}
	data, err := service.GetById(idx)
	if err != nil {
		c.String(http.StatusBadGateway, "Unable to  process data")

	}
	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return c.JSON(http.StatusOK, res)
}
