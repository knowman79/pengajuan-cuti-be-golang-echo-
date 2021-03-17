package repository

import (
	"example/driver"
	"example/models"
	"fmt"
	"time"
)

func ReadAllLeave() []models.LeaveModel {
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	defer db.Close()

	var result []models.LeaveModel

	items, err := db.Query(`select * from "tb_leave"`)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Printf("%T", items)

	for items.Next() {
		var each = models.LeaveModel{}
		var err = items.Scan(&each.FormId, &each.UserId, &each.Name, &each.Types, &each.CreatedBy, &each.ModifiedBy,
			&each.CreatedDate, &each.LastModifiedDate, &each.StartDate, &each.EndDate, &each.Description,
			&each.ReplacementId, &each.Address, &each.Phone, &each.Status)

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

func ReadIdLeave(userId int) []models.LeaveModel {
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	defer db.Close()

	var result []models.LeaveModel

	items, err := db.Query(`select * from "tb_leave" where user_id = $1`, userId)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Printf("%T", items)

	for items.Next() {
		var each = models.LeaveModel{}
		var err = items.Scan(&each.FormId, &each.UserId, &each.Name, &each.Types, &each.CreatedBy, &each.ModifiedBy,
			&each.CreatedDate, &each.LastModifiedDate, &each.StartDate, &each.EndDate, &each.Description,
			&each.ReplacementId, &each.Address, &each.Phone, &each.Status)

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

func CreateLeave(L *models.LeaveModel) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
		return Res
	}

	defer db.Close()

	date := time.Now()

	_, err = db.Exec(`INSERT INTO "tb_leave" ( "user_id", "name", "type", "created_by", "modified_by", "created_date", "last_modified_date",
					"start_date", "end_date", "description", "replacement_id", "address", "phone", "status")
					 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)`,
		L.UserId, L.Name, L.Types, L.CreatedBy, L.ModifiedBy, date, date, L.StartDate, L.EndDate, L.Description, L.ReplacementId, L.Address, L.Phone, L.Status)

	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed save Data"}
		return Res
	}
	fmt.Println("insert success!")
	Res = &ResponseModel{200, "Success save Data"}
	return Res
}

func DeleteLeave(formId int) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
		return Res
	}

	defer db.Close()

	_, err = db.Exec("delete from tb_leave where form_id = $1", formId)
	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed delete Data"}
		return Res
	}
	fmt.Println("Delete success!")
	Res = &ResponseModel{200, "Success Delete Data"}
	return Res
}

func UpdateLeave(L *models.LeaveModel, formId int) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())

		return Res
	}

	defer db.Close()
	date := time.Now()

	_, err = db.Exec("update tb_leave set user_id = $1, name = $2, type = $3 , created_by = $4, modified_by = $5, created_date = $6, last_modified_date =$7, start_date = $8, end_date = $9, description = $10, replacement_id =$11, address = $12, phone = $13, status = $14 where form_id = $15",
		L.UserId, L.Name, L.Types, L.CreatedBy, L.ModifiedBy, L.CreatedDate, date, L.StartDate, L.EndDate, L.Description, L.ReplacementId, L.Address, L.Phone, L.Status, formId)
	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed save Data"}
		return Res
	}
	fmt.Println("Update success!")
	Res = &ResponseModel{200, "Success save Data"}
	return Res
}

func DeleteLeaveDraft(formId int) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
		return Res
	}

	defer db.Close()

	_, err = db.Exec("delete from tb_leave where status='draft' and form_id = $1", formId)
	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed delete Data"}
		return Res
	}
	fmt.Println("Delete success!")
	Res = &ResponseModel{200, "Success Delete Data"}
	return Res
}

func UpdateLeaveApproved(L *models.LeaveModel, formId int) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())

		return Res
	}

	defer db.Close()

	_, err = db.Exec("update tb_leave set status = 'Approved' where status = 'Inprogress' and form_id = $1",
		formId)
	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed save Data"}
		return Res
	}
	fmt.Println("Update success!")
	Res = &ResponseModel{200, "Success save Data"}
	return Res
}

func UpdateLeaveOpenToInprogress(L *models.LeaveModel, formId int) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())

		return Res
	}

	defer db.Close()

	_, err = db.Exec("update tb_leave set status = 'Inprogress' where status = 'Open' and form_id = $1",
		formId)
	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed save Data"}
		return Res
	}
	fmt.Println("Update success!")
	Res = &ResponseModel{200, "Success save Data"}
	return Res
}
