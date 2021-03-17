package routes

import (
	"example/service"

	"github.com/labstack/echo"
)

//RoleEndPoint function
func Endpoint() {
	e := echo.New()
	//roles endpoint
	e.GET("/role/readAll", service.ReadAllRole)

	//user endpoint
	e.POST("/user/save", service.CreateUser)
	e.GET("/user/readAll", service.ReadAllUser)
	e.PATCH("/user/update", service.UpdateUser)
	e.DELETE("/user/delete", service.DeleteUser)

	//leave_allowance endpoint
	e.POST("/allowance/save", service.CreateAllowance)
	e.GET("/allowance/readAll", service.ReadAllAllowance)
	e.PATCH("/allowance/update", service.UpdateAllowance)
	e.DELETE("/allowance/delete", service.DeleteAllowance)
	
	e.Logger.Fatal(e.Start(":1323"))
}
