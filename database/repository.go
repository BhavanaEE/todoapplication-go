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

func CreateTodo(id any, title any, completed any) {
	db := InitDatabase()

	err := db.QueryRow("INSERT INTO todo(Id,Content, Completed) VALUES(?, ?, ?);", id, title, completed).Scan(&id)
	if err == nil {
		panic(err.Error())
	}
	defer db.Close()
}

func GetTodo(params map[string]string) *sql.Rows {
	db := InitDatabase()
	result, err := db.Query("SELECT Id, Content, Completed FROM todo WHERE Id = ?", params["Id"])
	if err != nil {
		panic(err.Error())
	}
	return result
}
