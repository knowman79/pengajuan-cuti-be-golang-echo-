package service

import (
	"example/models"
	"example/repository"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// // Create function
func CreateLeave(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	U := new(models.LeaveModel)
	if err := c.Bind(U); err != nil {
		return nil
	}
	Res = (*ResponseModel)(repository.CreateLeave(U))
	return c.JSON(http.StatusOK, Res)
}

// ReadAll function
func ReadAllLeave(c echo.Context) error {
	result := repository.ReadAllLeave()
	return c.JSON(http.StatusOK, result)
}

// ReadAll function
func ReadIdLeave(c echo.Context) error {
	id := c.QueryParam("userId")
	data, _ := strconv.Atoi(id)
	result := repository.ReadIdLeave(data)
	return c.JSON(http.StatusOK, result)
}

// DeleteUser function
func DeleteLeave(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	id := c.QueryParam("formId")
	data, _ := strconv.Atoi(id)
	Res = (*ResponseModel)(repository.DeleteLeave(data))
	return c.JSON(http.StatusOK, Res)
}

// UpdateUser function
func UpdateLeave(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	id := c.QueryParam("formId")
	data, _ := strconv.Atoi(id)
	U := new(models.LeaveModel)
	if err := c.Bind(U); err != nil {
		return nil
	}
	Res = (*ResponseModel)(repository.UpdateLeave(U, data))
	return c.JSON(http.StatusOK, Res)
}

func DeleteLeaveDraft(c echo.Context) error {
    Res := &ResponseModel{400, "Bad Request"}
    id := c.QueryParam("formId")
    data, _ := strconv.Atoi(id)
    Res = (*ResponseModel)(repository.DeleteLeaveDraft(data))
    return c.JSON(http.StatusOK, Res)
}