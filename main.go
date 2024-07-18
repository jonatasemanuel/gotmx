package main

import (
	"context"
	"net/http"

	"github.com/jonatasemanuel/gotmx-demo/templates"
	//"github.com/jonatasemanuel/gotmx-demo/templates/common"
	"github.com/jonatasemanuel/gotmx-demo/templates/components"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	// component.Render(context.Background(), os.Stdout)

	e.GET("/", func(c echo.Context) error {
		component := templates.Index()
		return component.Render(context.Background(), c.Response().Writer)
	})

	e.GET("/components", func(c echo.Context) error {
		typ := c.QueryParam("type")
		switch typ {
		case "add-todo":
			component := components.AddTodoInput()
			return component.Render(context.Background(), c.Response().Writer)
		case "add-todo-btn":
			component := components.AddTodoButton()
			return component.Render(context.Background(), c.Response().Writer)
		}

		return echo.NewHTTPError(http.StatusBadRequest, "Invalid element")
	})

	e.Static("/css", "css")
	e.Static("/fonts", "fonts")
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":8000"))
}
