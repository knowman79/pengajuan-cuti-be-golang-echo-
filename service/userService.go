package service

import (
	"example/models"
	"example/repository"
	"fmt"
	"strconv"

	"net/http"

	"github.com/labstack/echo"
)

// ResponseModel function
type ResponseModel struct {
	Code    int    `json:"code" validate:"required"`
	Message string `json:"message" validate:"required"`
}

// CreateUser function
func CreateUser(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	U := new(models.UserModel)
	if err := c.Bind(U); err != nil {
		return nil
	}
	Res = (*ResponseModel)(repository.CreateUser(U))
	return c.JSON(http.StatusOK, Res)
}

// ReadAllUser function
func ReadAllUser(c echo.Context) error {
	result := repository.ReadAllUser()
	return c.JSON(http.StatusOK, result)
}

// DeleteUser function
func DeleteUser(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	id := c.QueryParam("userId")
	data, _ := strconv.Atoi(id)
	fmt.Println("id", data)
	Res = (*ResponseModel)(repository.DeleteUser(data))
	return c.JSON(http.StatusOK, Res)
}

// UpdateUser function
func UpdateUser(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	id := c.QueryParam("userId")
	data, _ := strconv.Atoi(id)
	U := new(models.UserModel)
	if err := c.Bind(U); err != nil {
		return nil
	}
	Res = (*ResponseModel)(repository.UpdateUser(U, data))
	return c.JSON(http.StatusOK, Res)
}
