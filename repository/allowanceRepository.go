package repository

import (
	"example/driver"
	"example/models"
	"fmt"
)

type ResponseModelAllowance struct {
	Code    int    `json:"code" validate:"required"`
	Message string `json:"message" validate:"required"`
}

func ReadAllAllowance() []models.AllowanceModel {
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	defer db.Close()

	var result []models.AllowanceModel

	items, err := db.Query("select leave_id, user_id, current_leave, last_year_leave from leave_allowance")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Printf("%T", items)

	for items.Next() {
		var each = models.AllowanceModel{}
		var err = items.Scan(&each.LeaveId, &each.UserId, &each.CurrentLeave, &each.LastYearLeave)

		if err != nil {
			fmt.Println(err.Error())
			return nil
		}

		result = append(result, each)

	}

	if err = items.Err(); err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return result
}

func CreateAllowance(U *models.AllowanceModel) *ResponseModelAllowance {
	Res := &ResponseModelAllowance{500, "Internal Server Error"}
	db, err := driver.ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
		return Res
	}

	defer db.Close()

	_, err = db.Exec(`INSERT INTO "leave_allowance" ("leave_id", "user_id", "current_leave", "last_year_leave") VALUES ($1, $2, $3, $4)`, U.LeaveId, U.UserId, U.CurrentLeave, U.LastYearLeave)

	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModelAllowance{400, "Failed save Data"}
		return Res
	}
	fmt.Println("insert success!")
	Res = &ResponseModelAllowance{200, "Success save Data"}
	return Res
}

func DeleteAllowance(leaveId int) *ResponseModelAllowance {
	Res := &ResponseModelAllowance{500, "Internal Server Error"}
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
		return Res
	}

	defer db.Close()

	_, err = db.Exec("delete from leave_allowance where leave_id = $1", leaveId)
	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModelAllowance{400, "Failed save Data"}
		return Res
	}
	fmt.Println("Delete success!")
	Res = &ResponseModelAllowance{200, "Success save Data"}
	return Res
}

func UpdateAllowance(U *models.AllowanceModel, leaveId int) *ResponseModelAllowance {
	Res := &ResponseModelAllowance{500, "Internal Server Error"}
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())

		return Res
	}

	defer db.Close()

	_, err = db.Exec("update leave_allowance set user_id = $1, current_leave = $2, last_year_leave = $3 where leave_id = $5", U.UserId, U.CurrentLeave, U.LastYearLeave, leaveId)
	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModelAllowance{400, "Failed save Data"}
		return Res
	}
	fmt.Println("Update success!")
	Res = &ResponseModelAllowance{200, "Success save Data"}
	return Res
}
