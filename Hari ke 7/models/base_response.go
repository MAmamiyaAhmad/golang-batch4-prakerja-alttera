package routes

import (
	"prakerja4/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.GET("/users", controllers.GetUserController)
	e.POST("/users", controllers.InsertUserController)
	return e
}