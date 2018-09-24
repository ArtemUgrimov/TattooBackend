package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func restoreHandler(c echo.Context) error {
	userid := c.FormValue("id")
	if sess, ok := sessions[userid]; ok {
		return c.File(sess.backupName)
	}
	return c.JSON(http.StatusNotFound, map[string]string{
		"status": "error",
		"cause":  "session not found",
	})
}
