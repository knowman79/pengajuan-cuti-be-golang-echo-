package repository

import (
	"example/driver"
	"example/models"
	"fmt"
	"time"
)

func ReadAllLeave() []models.ResponseLeaveModel {
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	defer db.Close()

	var result []models.ResponseLeaveModel

	items, err := db.Query(`select form_id, tb_leave.user_id, tb_leave.name, type, created_by, modified_by, created_date, last_modified_date, start_date, end_date, end_date-start_date, description, replacement_id, address, phone, status, tb_leave.leave_id, tb_user."name", tb_leave_allowance."current_leave" from tb_leave JOIN tb_user ON replacement_id = tb_user.user_id JOIN tb_leave_allowance ON tb_leave.user_id = tb_leave_allowance.user_id`)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Printf("%T", items)

	for items.Next() {
		var each = models.ResponseLeaveModel{}
		var err = items.Scan(&each.FormId, &each.UserId, &each.Name, &each.Types, &each.CreatedBy, &each.ModifiedBy,
			&each.CreatedDate, &each.LastModifiedDate, &each.StartDate, &each.EndDate, &each.LeaveDuration, &each.Description,
			&each.ReplacementId, &each.Address, &each.Phone, &each.Status, &each.LeaveId, &each.ReplacementName, &each.Current_leave)

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
			&each.ReplacementId, &each.Address, &each.Phone, &each.Status, &each.LeaveId)

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

func CreateLeave(L *models.LeaveModel, M *models.AllowanceModel) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}

	db, err := driver.ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
		return Res
	}

	defer db.Close()

	date := time.Now()
	duration := L.EndDate.Day() - L.StartDate.Day()

	_, err = db.Exec(`WITH shape_select as (
                        INSERT INTO "tb_leave" ("user_id", "name", "type", "created_by", "modified_by", "created_date", "last_modified_date",
                        "start_date", "end_date", "description", "replacement_id", "address", "phone", "status", "leave_id", "duration")
                        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
                        returning leave_id, user_id, duration)
                    SELECT leave_id
                    from tb_leave_allowance
                    where tb_leave_allowance.leave_id = (select leave_id from shape_select) and (select duration from shape_select) < (tb_leave_allowance.current_leave+tb_leave_allowance.last_year_leave)`,
		L.UserId, L.Name, L.Types, L.CreatedBy, L.ModifiedBy, date, date, L.StartDate, L.EndDate, L.Description, L.ReplacementId, L.Address, L.Phone, L.Status, L.LeaveId, duration)

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

func UpdateLeave(L *models.LeaveModel) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())

		return Res
	}

	defer db.Close()
	date := time.Now()

	_, err = db.Exec("update tb_leave set user_id = $1, name = $2, type = $3 , created_by = $4, modified_by = $5, created_date = $6, last_modified_date =$7, start_date = $8, end_date = $9, description = $10, replacement_id =$11, address = $12, phone = $13, status = $14, leave_id = $15 where form_id = $16",
		L.UserId, L.Name, L.Types, L.CreatedBy, L.ModifiedBy, L.CreatedDate, date, L.StartDate, L.EndDate, L.Description, L.ReplacementId, L.Address, L.Phone, L.Status, L.LeaveId, L.FormId)
	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed save Data"}
		return Res
	}
	fmt.Println("Update success!")
	Res = &ResponseModel{200, "Success save Data"}
	return Res
}

func DeleteLeaveDraft(L *models.LeaveModel) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
		return Res
	}

	defer db.Close()

	_, err = db.Exec("delete from tb_leave where status='Draft' and form_id = $1", L.FormId)
	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed delete Data"}
		return Res
	}
	fmt.Println("Delete success!")
	Res = &ResponseModel{200, "Success Delete Data"}
	return Res
}

func UpdateLeaveApproved(L *models.LeaveModel) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())

		return Res
	}

	defer db.Close()

	_, err = db.Exec("update tb_leave set status = 'Approved' where status = 'Inprogress' and form_id = $1",
		L.FormId)
	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed save Data"}
		return Res
	}
	fmt.Println("Update success!")
	Res = &ResponseModel{200, "Success save Data"}
	return Res
}

func UpdateLeaveOpenToInprogress(L *models.LeaveModel) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())

		return Res
	}

	defer db.Close()

	_, err = db.Exec("update tb_leave set status = 'Inprogress' where status = 'Open' and form_id = $1",
		L.FormId)
	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed save Data"}
		return Res
	}
	fmt.Println("Update success!")
	Res = &ResponseModel{200, "Success save Data"}
	return Res
}

func ReadLeaveByName(name string) []models.LeaveByNameModel {
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	defer db.Close()

	var result []models.LeaveByNameModel

	items, err := db.Query(`SELECT created_date, start_date, end_date, type , status, duration FROM tb_leave WHERE name =$1 `, name)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	fmt.Printf("%T\n", items)

	for items.Next() {
		var each = models.LeaveByNameModel{}
		var err = items.Scan(&each.CreatedDate, &each.StartDate, &each.EndDate, &each.Type,
			&each.Status, &each.Duration)

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

func UpdateLeaveDraftToOpen(L *models.LeaveModel, M *models.AllowanceModel) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())

		return Res
	}

	defer db.Close()

	var leaveTotal int
	leaveTotal = (L.EndDate.Day() - L.StartDate.Day())

	if M.LastYearLeave > leaveTotal {
		_, err = db.Exec("with shape_update as ( UPDATE tb_leave SET status = 'Open' WHERE tb_leave.form_id = $1 and status = 'Draft' returning tb_leave.leave_id, tb_leave.duration) UPDATE tb_leave_allowance SET last_year_leave = last_year_leave - (select duration from shape_update) WHERE (tb_leave_allowance.leave_id) IN (select leave_id from shape_update)",
			L.FormId)

	} else {
		_, err = db.Exec("with shape_update as ( UPDATE tb_leave SET status = 'Open' WHERE tb_leave.form_id = $1 and status = 'Draft' returning tb_leave.leave_id, tb_leave.duration) UPDATE tb_leave_allowance SET last_year_leave = 0, current_leave = current_leave - ((select duration from shape_update) - last_year_leave) WHERE (tb_leave_allowance.leave_id) IN (select leave_id from shape_update);",
			L.FormId)
	}

	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed save Data"}
		return Res
	}

	fmt.Println("Update success!")
	Res = &ResponseModel{200, "Success save Data"}
	return Res
}

func UpdateLeaveCanceled(L *models.LeaveModel, M *models.AllowanceModel) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())

		return Res
	}

	defer db.Close()

	//var comebackLeave int = (L.Duration + M.CurrentLeave)

	//if comebackLeave > 12 {
	_, err = db.Exec("with shape_update as ( UPDATE tb_leave SET status = 'Canceled' WHERE tb_leave.form_id = $1 and status = 'Open' returning tb_leave.leave_id, tb_leave.duration) UPDATE tb_leave_allowance SET current_leave = 12, last_year_leave = last_year_leave + (((select duration from shape_update) + current_leave) - 12) WHERE ((select duration from shape_update) + current_leave) > 12 and (tb_leave_allowance.leave_id) IN (select leave_id from shape_update)", L.FormId)

	//} else {
	_, err = db.Exec("with shape_update as ( UPDATE tb_leave SET status = 'Canceled' WHERE tb_leave.form_id = $1 and status = 'Open' returning tb_leave.leave_id, tb_leave.duration) UPDATE tb_leave_allowance SET current_leave = current_leave + (select duration from shape_update) WHERE ((select duration from shape_update) + current_leave) < 13 and (tb_leave_allowance.leave_id) IN (select leave_id from shape_update)", L.FormId)
	//}

	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed save Data"}
		return Res
	}

	fmt.Println("Update success!")
	Res = &ResponseModel{200, "Success save Data"}
	return Res
}

func UpdateRejectBySPV(L *models.LeaveModel, M *models.AllowanceModel) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())

		return Res
	}

	defer db.Close()

	_, err = db.Exec("with shape_update as ( UPDATE tb_leave SET status = 'Reject by Supervisor' WHERE tb_leave.form_id = $1 and status = 'Open' returning tb_leave.leave_id, tb_leave.duration) UPDATE tb_leave_allowance SET current_leave = 12, last_year_leave = last_year_leave + (((select duration from shape_update) + current_leave) - 12) WHERE ((select duration from shape_update) + current_leave) > 12 and (tb_leave_allowance.leave_id) IN (select leave_id from shape_update)", L.FormId)

	_, err = db.Exec("with shape_update as ( UPDATE tb_leave SET status = 'Reject by Supervisor' WHERE tb_leave.form_id = $1 and status = 'Open' returning tb_leave.leave_id, tb_leave.duration) UPDATE tb_leave_allowance SET current_leave = current_leave + (select duration from shape_update) WHERE ((select duration from shape_update) + current_leave) < 13 and (tb_leave_allowance.leave_id) IN (select leave_id from shape_update)", L.FormId)

	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed save Data"}
		return Res
	}

	fmt.Println("Update success!")
	Res = &ResponseModel{200, "Success save Data"}
	return Res

}

func UpdateLeaveRejectByHRD(L *models.LeaveModel, M *models.AllowanceModel) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())

		return Res
	}

	defer db.Close()

	_, err = db.Exec("with shape_update as ( UPDATE tb_leave SET status = 'Reject by HRD' WHERE tb_leave.form_id = $1 and status = 'Inprogress' returning tb_leave.leave_id, tb_leave.duration) UPDATE tb_leave_allowance SET current_leave = 12, last_year_leave = last_year_leave + (((select duration from shape_update) + current_leave) - 12) WHERE ((select duration from shape_update) + current_leave) > 12 and (tb_leave_allowance.leave_id) IN (select leave_id from shape_update)", L.FormId)

	_, err = db.Exec("with shape_update as ( UPDATE tb_leave SET status = 'Reject by HRD' WHERE tb_leave.form_id = $1 and status = 'Inprogress' returning tb_leave.leave_id, tb_leave.duration) UPDATE tb_leave_allowance SET current_leave = current_leave + (select duration from shape_update) WHERE ((select duration from shape_update) + current_leave) < 13 and (tb_leave_allowance.leave_id) IN (select leave_id from shape_update)", L.FormId)

	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed save Data"}
		return Res
	}

	fmt.Println("Update success!")
	Res = &ResponseModel{200, "Success save Data"}
	return Res

}

func UpdateStatusDraft(L *models.LeaveModel, M *models.AllowanceModel) *ResponseModel {
	Res := &ResponseModel{500, "Internal Server Error"}
	db, err := driver.ConnectDB()

	if err != nil {
		fmt.Println(err.Error())

		return Res
	}

	defer db.Close()

	ed := time.Date(L.EndDate.Year(), L.EndDate.Month(), L.EndDate.Day(), 0, 0, 0, 0, time.UTC)
	sd := time.Date(L.StartDate.Year(), L.StartDate.Month(), L.StartDate.Day(), 0, 0, 0, 0, time.UTC)
	du := ed.Sub(sd).Hours() / 24

	_, err = db.Exec("with shape_update as ( SELECT tb_leave_allowance.leave_id, current_leave, last_year_leave from tb_leave_allowance join tb_leave on tb_leave_allowance.leave_id = tb_leave.leave_id where tb_leave.form_id = $10) UPDATE tb_leave SET type = $1, start_date = $2, end_date = $3, description = $4, replacement_id = $5, address = $6, phone = $7, duration = $8 WHERE status = 'Draft' and form_id = $9 and (select leave_id from shape_update)= tb_leave.leave_id and $8 < ((select current_leave from shape_update) + (select last_year_leave from shape_update))",
		L.Types, L.StartDate, L.EndDate, L.Description, L.ReplacementId, L.Address, L.Phone, du, L.FormId, L.FormId)
	if err != nil {
		fmt.Println(err.Error())
		Res = &ResponseModel{400, "Failed save Data"}
		return Res
	}
	fmt.Println("Update success!")
	Res = &ResponseModel{200, "Success save Data"}
	return Res
}
