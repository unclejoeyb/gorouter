package main

import (
	"context"
	"os"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	component := hello.hello("John")
	component.Render(context.Background(), os.Stdout)
	e.GET("/", func(c echo.Context) error {
		return component.Render(c.Response().Writer)
	})
	e.Static("/", "static")
	e.Logger.Fatal(e.Start(":3000"))
	
}






