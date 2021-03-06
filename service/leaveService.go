package service

import (
	"example/models"
	"example/repository"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// CreateLeave . . .
func CreateLeave(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	U := new(models.LeaveModel)
	L := new(models.AllowanceModel)
	if err := c.Bind(U); err != nil {
		log.Println(err.Error())
		return nil
	}
	Res = (*ResponseModel)(repository.CreateLeave(U, L))
	return c.JSON(http.StatusOK, Res)
}

// ReadAllLeave . . .
func ReadAllLeave(c echo.Context) error {
	result := repository.ReadAllLeave()
	return c.JSON(http.StatusOK, result)
}

// ReadIdLeave . . .
func ReadIDLeave(c echo.Context) error {
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
	U := new(models.LeaveModel)
	if err := c.Bind(U); err != nil {
		log.Println(err.Error())
		return nil
	}
	Res = (*ResponseModel)(repository.UpdateLeave(U))
	return c.JSON(http.StatusOK, Res)
}

func DeleteLeaveDraft(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	U := new(models.LeaveModel)
	Res = (*ResponseModel)(repository.DeleteLeaveDraft(U))
	return c.JSON(http.StatusOK, Res)
}

func UpdateLeaveApproved(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	U := new(models.LeaveModel)
	if err := c.Bind(U); err != nil {
		log.Println(err.Error())
		return nil
	}
	Res = (*ResponseModel)(repository.UpdateLeaveApproved(U))
	return c.JSON(http.StatusOK, Res)
}

func UpdateLeaveOpenToInprogress(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	U := new(models.LeaveModel)
	if err := c.Bind(U); err != nil {
		log.Println(err.Error())
		return nil
	}
	Res = (*ResponseModel)(repository.UpdateLeaveOpenToInprogress(U))
	return c.JSON(http.StatusOK, Res)
}

func ReadLeaveByName(c echo.Context) error {
	name := c.QueryParam("name")
	result := repository.ReadLeaveByName(name)
	return c.JSON(http.StatusOK, result)
}

func UpdateLeaveDraftToOpen(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	U := new(models.LeaveModel)
	L := new(models.AllowanceModel)
	if err := c.Bind(U); err != nil {
		log.Println(err.Error())
		return nil
	}
	Res = (*ResponseModel)(repository.UpdateLeaveDraftToOpen(U, L))
	return c.JSON(http.StatusOK, Res)
}

func UpdateLeaveCanceled(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	U := new(models.LeaveModel)
	L := new(models.AllowanceModel)
	if err := c.Bind(U); err != nil {
		log.Println(err.Error())
		return nil
	}
	Res = (*ResponseModel)(repository.UpdateLeaveCanceled(U, L))
	return c.JSON(http.StatusOK, Res)
}

func UpdateRejectBySPV(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	U := new(models.LeaveModel)
	L := new(models.AllowanceModel)
	if err := c.Bind(U); err != nil {
		log.Println(err.Error())
		return nil
	}
	Res = (*ResponseModel)(repository.UpdateLeaveRejectByHRD(U, L))
	return c.JSON(http.StatusOK, Res)
}

func UpdateLeaveRejectByHRD(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	U := new(models.LeaveModel)
	L := new(models.AllowanceModel)
	if err := c.Bind(U); err != nil {
		log.Println(err.Error())
		return nil
	}
	Res = (*ResponseModel)(repository.UpdateLeaveRejectByHRD(U, L))
	return c.JSON(http.StatusOK, Res)
}

func UpdateStatusDraft(c echo.Context) error {
	Res := &ResponseModel{400, "Bad Request"}
	U := new(models.LeaveModel)
	L := new(models.AllowanceModel)
	if err := c.Bind(U); err != nil {
		log.Println(err.Error())
		return nil
	}
	Res = (*ResponseModel)(repository.UpdateStatusDraft(U, L))
	return c.JSON(http.StatusOK, Res)
}
