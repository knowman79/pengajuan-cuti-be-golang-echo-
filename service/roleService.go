package service

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/my/repo/repository"
)

func ReadAllRole(c echo.Context) error {
	result := repository.ReadAll()
	return c.JSON(http.StatusOK, result)
}
