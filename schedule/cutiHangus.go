package schedule

import (
	"example/driver"
	"log"
)

func CutiHangusScheduler() {
	db, err := driver.ConnectDB()

	log.Println("Running Scheduler")

	if err != nil {
		log.Println(err.Error())
	}

	sqlStatementUpdate := `UPDATE tb_leave_allowance set last_year_leave = 0`
	_, e := db.Exec(sqlStatementUpdate)
	if e != nil {
		log.Println(err.Error())
	}
}
