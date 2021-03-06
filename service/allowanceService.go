package service

import (
	"example/models"
	"example/repository"
	"fmt"
	"log"
	"strconv"

	"net/http"

	"github.com/labstack/echo"
)

// ResponseModelAllowance . . .
type ResponseModelAllowance struct {
	Code    int    `json:"code" validate:"required"`
	Message string `json:"message" validate:"required"`
}

// CreateAllowance . . .
func CreateAllowance(c echo.Context) error {
	Res := &ResponseModelAllowance{400, "Bad Request"}
	U := new(models.AllowanceModel)
	if err := c.Bind(U); err != nil {
		log.Println(err.Error())
		return nil
	}
	Res = (*ResponseModelAllowance)(repository.CreateAllowance(U))
	return c.JSON(http.StatusOK, Res)
}

// ReadAllAllowance . . .
func ReadAllAllowance(c echo.Context) error {
	result := repository.ReadAllAllowance()
	return c.JSON(http.StatusOK, result)
}

// DeleteAllowance . . .
func DeleteAllowance(c echo.Context) error {
	Res := &ResponseModelAllowance{400, "Bad Request"}
	id := c.QueryParam("leaveId")
	data, _ := strconv.Atoi(id)
	fmt.Println("id", data)
	Res = (*ResponseModelAllowance)(repository.DeleteAllowance(data))
	return c.JSON(http.StatusOK, Res)
}

// UpdateAllowance . . .
func UpdateAllowance(c echo.Context) error {
	Res := &ResponseModelAllowance{400, "Bad Request"}
	id := c.QueryParam("leaveId")
	data, _ := strconv.Atoi(id)
	U := new(models.AllowanceModel)
	if err := c.Bind(U); err != nil {
		log.Println(err.Error())
		return nil
	}
	Res = (*ResponseModelAllowance)(repository.UpdateAllowance(U, data))
	return c.JSON(http.StatusOK, Res)
}
