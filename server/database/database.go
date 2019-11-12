package database

import (
	"backend-test/config"
	"backend-test/utils"
	"database/sql"
	"log"

	//Importation Postgres connection driver
	_ "github.com/lib/pq"
)

//Conn is a global connect to database
var Conn *sql.DB

//Connect opens connection to postgres database
func Connect() {
	connectionString := utils.BuildString(`postgres://`, config.DbUser, `:`, config.DbPassword, `@`, config.DbHost, `/`, config.DbName, `?sslmode=disable`)
	conn, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Println(err.Error())
		return
	}

	Conn = conn
}
