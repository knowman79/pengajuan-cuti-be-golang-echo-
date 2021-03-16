package routes

import (
	"github.com/labstack/echo"
	"github.com/my/repo/service"
)

func RoleEndpoint() {
	e := echo.New()
	e.GET("/role/readAll", service.ReadAllRole)

	e.Logger.Fatal(e.Start(":1323"))
}
