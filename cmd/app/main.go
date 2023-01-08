package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/louisfield/go-app-backend/cmd/session"
)

// Gen IDs with UUIDS
// Have user module to store user rather than private module.

func main() {
	// Hello world, the web server

	e := echo.New()

	// Middleware

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/session", session.CreateSession)
	e.POST("/session/:id", session.AddUserToSession)

	e.Logger.Fatal(e.Start(":8000"))
}
