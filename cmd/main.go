package main

import (
	"github.com/jmarren/htmx-form-example/internal/render"
	"github.com/jmarren/htmx-form-example/internal/routers"
	"github.com/labstack/echo/v4"
)

// Initialize Echo framework and templates
func initEcho() *echo.Echo {
	render.InitTemplates()

	// Create a template registry for Echo
	e := echo.New()
	e.Renderer = render.TmplRegistry

	// Serve static files
	e.Static("/static", "ui/static")

	return e
}

func main() {
	e := initEcho()

	e.GET("/sign-in", routers.GetSignInForm)
	e.POST("/sign-in", routers.SignIn)

	e.Logger.Fatal(e.Start(":3000"))
}
