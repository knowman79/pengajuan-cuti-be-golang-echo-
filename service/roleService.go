package service

import (
	"example/repository"

	"net/http"

	"github.com/labstack/echo"
)

func ReadAllRole(c echo.Context) error {
	result := repository.ReadAll()
	return c.JSON(http.StatusOK, result)
}
