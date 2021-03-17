package routes

import (
	"example/service"

	"github.com/labstack/echo"
)

//UserEndPoint function
func UserEndpoint() {
	e := echo.New()
	e.POST("/user/save", service.CreateUser)
	e.GET("/user/readAll", service.ReadAllUser)
	e.PATCH("/user/update", service.UpdateUser)
	e.DELETE("/user/delete", service.DeleteUser)
	e.Logger.Fatal(e.Start(":1323"))
}
