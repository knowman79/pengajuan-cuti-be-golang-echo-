package routes

import (
	"example/service"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	//"github.com/labstack/echo/v4/middleware"
)

//RoleEndPoint function
func Endpoint() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// CORS restricted
	// Allows requests from any `https://labstack.com` or `https://labstack.net` origin
	// wth GET, PUT, POST or DELETE method.
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	//roles endpoint
	e.GET("/role/readAll", service.ReadAllRole)

	//user endpoint
	e.POST("/user/save", service.CreateUser)
	e.GET("/user/readAll", service.ReadAllUser)
	e.PATCH("/user/update", service.UpdateUser)
	e.DELETE("/user/delete", service.DeleteUser)
	e.GET("/user/list", service.ReadAllOlUser)

	//leave_allowance endpoint
	e.POST("/allowance/save", service.CreateAllowance)
	e.GET("/allowance/readAll", service.ReadAllAllowance)
	e.PATCH("/allowance/update", service.UpdateAllowance)
	e.DELETE("/allowance/delete", service.DeleteAllowance)

	//leave endpoint
	e.GET("/leave/readAll", service.ReadAllLeave)
	e.POST("/leave/save", service.CreateLeave)
	e.DELETE("/leave/delete", service.DeleteLeave)
	e.PATCH("/leave/update", service.UpdateLeave)
	e.GET("/leave/readById", service.ReadIdLeave)
	e.DELETE("/leave/deleteDraft", service.DeleteLeaveDraft)
	e.PATCH("/leave/updateApproved", service.UpdateLeaveApproved)
	e.PATCH("/leave/updateToInprogress", service.UpdateLeaveOpenToInprogress)

	e.Logger.Fatal(e.Start(":1323"))
}
