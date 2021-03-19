package schedule

import (
	"example/driver"
	"example/models"
	"log"
)

func CrossYearLeaveScheduler() {
	db, err := driver.ConnectDB()

	if err != nil {
		log.Println(err.Error())
	}

	sqlStatementSelect := `SELECT * FROM tb_leave_allowance`

	items, err := db.Query(sqlStatementSelect)
	if err != nil {
		log.Println(err.Error())
	}

	defer items.Close()
	for items.Next() {
		var each = models.AllowanceModel{}
		var userId int
		var err = items.Scan(&each.LeaveId, &userId, &each.CurrentLeave, &each.LastYearLeave)

		if err != nil {
			log.Println(err.Error())
		}

		each.LastYearLeave = each.CurrentLeave
		each.CurrentLeave = 12

		sqlStatementUpdate := `UPDATE tb_leave_allowance set last_year_leave = $1, current_leave= $2  WHERE leave_id = $3`
		_, e := db.Exec(sqlStatementUpdate, each.LastYearLeave, each.CurrentLeave, each.LeaveId)
		if e != nil {
			log.Println(err.Error())
		}
	}

	if err = items.Err(); err != nil {
		log.Println(err.Error())
	}

}
