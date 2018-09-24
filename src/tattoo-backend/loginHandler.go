package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func loginHandler(c echo.Context) error {
	var status string

	userid := c.FormValue("id")
	if _, ok := sessions[userid]; ok {
		status = "reconnect"
	} else {
		session := Session{
			userid,
			userid + ".db",
		}
		sessions[userid] = session
		status = "new"
	}

	return c.JSON(http.StatusOK, map[string]string{"status": status})
}
