package app

import (
	"database/sql"
	"testing"
	"todos/helper"

	_ "github.com/go-sql-driver/mysql"
)

func TestConnection(t *testing.T) {
	db, err := sql.Open("mysql", "root@tcp(localhost:4343:3306)/journey_of_learn_golang_todos")
	helper.PanicIfError(err)

	defer db.Close()
}
