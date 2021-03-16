package delivery

import (
	"example/service"

	"github.com/labstack/echo"
)

func MakananEndPoint() {
	e := echo.New()
	e.GET("/v1/api/makanan/readAll", service.ReadAllMakanan)

	e.Logger.Fatal(e.Start(":1323"))
}
