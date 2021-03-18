package schedule

import (
	"example/driver"
	"fmt"
)

func ExpiredLeaveScheduler() {
	db, err := driver.ConnectDB()
	if err != nil {
		fmt.Println(err.Error())
	}

	sqlStatementUpdate := `UPDATE tb_leave_allowance set last_year_leave = 0`
	_, e := db.Exec(sqlStatementUpdate)
	if e != nil {
		fmt.Println(err.Error())
	}
}
