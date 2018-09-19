package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.POST("/backup", backupHandler)
	e.File("/restore", "db_backup.db")

	e.Logger.Fatal(e.Start(":8000"))
}
