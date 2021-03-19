package driver

import (
	"database/sql"
	"fmt" // package used to read the .env file
	"log"

	_ "github.com/lib/pq" // postgres golang driver
)

const (
	user     = "wrdev"
	dbname   = "dbcuti_m3"
	password = "Majuber54Ma2020!"
	host     = "139.99.67.15"
	port     = 5423
)

/*const (
	user     = "postgres"
	dbname   = "cuti"
	password = "1234"
	host     = "localhost"
	port     = 5432
)*/

//ConnectDB function
func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}

	fmt.Println("Successfully connected!")
	return db, nil
}
