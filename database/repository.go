package database

import (
	"database/sql"
)

func GetAllTodos() *sql.Rows {
	db := InitDatabase()
	result, err := db.Query("SELECT * from todo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	return result
}
