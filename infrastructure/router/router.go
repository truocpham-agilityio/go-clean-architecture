package router

import (
	"go-clean-architecture/interface/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", func(ctx echo.Context) error {
		return c.GetUsers(ctx)
	})

	return e
}
