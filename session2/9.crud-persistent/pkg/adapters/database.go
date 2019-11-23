package adapters

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// NewDbConnection initialize and returns a new database connection
func NewDbConnection() *sql.DB {
	db, err := sql.Open("mysql", "devfest_go:goturns10@tcp(35.213.163.65:3306)/phonebook")

	if err != nil {
		panic(err)
	}

	// Make sure that we have a proper connection with the server
	err = db.Ping()
	if  err != nil {
		panic(err)
	}

	return db
}
