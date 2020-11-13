package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var Connection *sql.DB

func Initialize(user, password, dbname string) *sql.DB {

	if Connection == nil {
		connectionString :=
			fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

		handle, err := sql.Open("postgres", connectionString)
		if err != nil {
			log.Fatal(err)
		}
		Connection = handle
	}
	return Connection
}
