package main

import (
	"context"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/unclejoeyb/gorouter/templates"
)

func main() {
	e := echo.New()
	component := templates.Index()
	component.Render(context.Background(), os.Stdout)
	e.GET("/", func(c echo.Context) error {
		return component.Render(context.Background(), c.Response().Writer)
	})
	e.Static("/static", "static")
	e.Static("/css", "css")
	e.Logger.Fatal(e.Start(":3000"))
	
}






