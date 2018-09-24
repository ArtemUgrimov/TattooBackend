package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func statusHandler(c echo.Context) error {
	var status string
	userid := c.FormValue("id")
	if _, ok := sessions[userid]; ok {
		status = "loggedIn"
	} else {
		status = "empty"
	}

	return c.JSON(http.StatusOK, map[string]string{"status": status})
}
