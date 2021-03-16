package delivery

import (
	"github.com/labstack/echo"
	"github.com/my/repo/service"
)

func UserEndpoint() {
	e := echo.New()
	e.GET("/v1/api/user/readAll", service.ReadAllUsers)
	e.POST("/v1/api/user/save", service.SaveUser)
	e.POST("/v1/api/user/update", service.UpdateUser)
	e.POST("/v1/api/user/delete", service.DeleteUser)

	e.GET("/v1/api/makanan/readAll", service.ReadAllMakanan)
	e.POST("/v1/api/makanan/save", service.SaveMakanan)
	e.POST("/v1/api/makanan/update", service.UpdateMakanan)
	e.POST("/v1/api/makanan/delete", service.DeleteMakanan)

	e.Logger.Fatal(e.Start(":1323"))
}
