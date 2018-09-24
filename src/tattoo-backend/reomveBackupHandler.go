package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func reomveBackupHandler(c echo.Context) error {
	status := "ok"

	return c.JSON(http.StatusOK, map[string]string{"status": status})
}
