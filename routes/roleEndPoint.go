package routes

import (
	"example/service"

	"github.com/labstack/echo"
)

func RoleEndpoint() {
	e := echo.New()
	e.GET("/role/readAll", service.ReadAllRole)

	e.Logger.Fatal(e.Start(":1323"))
}
