package delivery

import (
	"github.com/labstack/echo"
	"github.com/my/repo/service"
)

func MakananEndPoint() {
	e := echo.New()
	e.GET("/v1/api/makanan/readAll", service.ReadAllMakanan)

	e.Logger.Fatal(e.Start(":1323"))
}
