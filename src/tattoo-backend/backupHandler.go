package main

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo"
)

func backupHandler(c echo.Context) error {

	userid := c.FormValue("id")
	if userid == "" {
		return c.JSON(http.StatusOK, map[string]string{
			"status": "error",
			"cause":  "id not defined",
		})
	}

	if session, ok := sessions[userid]; ok {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		dst, err := os.Create(session.backupName)
		if err != nil {
			return err
		}
		defer dst.Close()

		if _, err = io.Copy(dst, src); err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]string{"status": "ok"})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"status": "error",
		"cause":  "not logged in",
	})
}
