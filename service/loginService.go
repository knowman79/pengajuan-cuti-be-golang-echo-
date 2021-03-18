package service

import (
	"example/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

var APPLICATION_NAME = "Pengajuan Cuti"
var LOGIN_EXPIRATION_DURATION = time.Duration(24) * time.Hour
var JWT_SIGNING_METHOD = jwt.SigningMethodHS256
var JWT_SIGNATURE_KEY = []byte("bematrix1")

func Login(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	U := new(models.LoginModel)
	if err := c.Bind(U); err != nil {
		return nil
	}
	//Res = (*ResponseModel)(repository.CreateLeave(U))
	return c.JSON(http.StatusOK, Res)
}
