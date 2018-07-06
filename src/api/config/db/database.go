package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"quoter/src/api/config/loggers"
)

var Database *sql.DB

const (
	host = "localhost"
	port = 5432
	password = "quoter"
	user = "quoter"
	dbName = "quoter"
)

func ConnectAndSetDatabase() {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host,
		port,
		user,
		password,
		dbName)

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		loggers.Error.Println("Couldn't establish connection to the database.", err)
		panic(err)
	}

	err = db.Ping()  // Needed to actually open a connection.
	if err != nil {
		loggers.Error.Println("An error occurred trying to ping the database.", err)
		panic(err)
	}

	loggers.Info.Println("Successfully connected!")
	Database = db
}
