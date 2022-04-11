package database

import (
	"database/sql"
	"todoapplication/model"
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

func GetTodo(params string) *sql.Rows {
	db := InitDatabase()
	result, err := db.Query("SELECT Id, Content, Completed FROM todo WHERE Id = ?", params)
	if err != nil {
		panic(err.Error())
	}
	return result
}

func UpdateTodo(params model.Todo) {
	db := InitDatabase()
	stmt, err := db.Prepare("UPDATE todo SET Content = ?, Completed=? WHERE Id = ?")
	_, err = stmt.Exec(params.Content, params.Completed, params.Id)
	if err != nil {
		panic(err.Error())
	}
}

func DeleteTodo(param string) {
	db := InitDatabase()
	stmt, err := db.Prepare("DELETE FROM todo WHERE Id = ?")
	_, err = stmt.Exec(param)
	if err != nil {
		panic(err.Error())
	}
}
