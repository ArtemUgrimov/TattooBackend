package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/status", statusHandler)
	e.POST("/login", loginHandler)
	e.POST("/logout", logoutHandler)
	e.POST("/backup", backupHandler)
	e.POST("/remove_backup", reomveBackupHandler)
	e.GET("/restore", restoreHandler)

	e.Logger.Fatal(e.Start(":8000"))
}
