package service

import (
	"example/repository"

	"net/http"

	"github.com/labstack/echo"
)

//ReadAllRole function
func ReadAllRole(c echo.Context) error {
	result := repository.ReadAllRole()
	return c.JSON(http.StatusOK, result)
}
