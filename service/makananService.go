package service

import (
	"example/model"
	"example/repository"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

func ReadAllMakanan(c echo.Context) error {
	result := repository.ReadAllMakanan()
	log.Println("Success Read All Makanan ", result)
	return c.JSON(http.StatusOK, result)
}

func SaveMakanan(c echo.Context) error {
	makanan := new(model.MakananModel)
	if err := c.Bind(makanan); err != nil {
		return nil
	}
	result := repository.SaveMakanan(makanan)
	log.Println("Success Save Makanan ", makanan)
	return c.JSON(http.StatusOK, result)
}

func UpdateMakanan(c echo.Context) error {
	makanan := new(model.MakananModel)
	if err := c.Bind(makanan); err != nil {
		return nil
	}
	result := repository.UpdateMakanan(makanan)
	log.Println("Success Upate Makanan ", makanan)
	return c.JSON(http.StatusOK, result)
}
func DeleteMakanan(c echo.Context) error {
	Res := model.ResponseModel{400, "Bad Request"}
	id := c.QueryParam("id")
	data, _ := strconv.Atoi(id)
	fmt.Println("id", data)
	Res = repository.DeleteMakanan(data)
	return c.JSON(http.StatusOK, Res)
}
