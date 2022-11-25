package app_test

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/journey_of_learn_golang_todos")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}
