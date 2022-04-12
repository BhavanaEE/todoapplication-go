package database

import (
	"database/sql"
	"todoapplication/model"
)

func GetAllTodos() (*sql.Rows, error) {
	db := InitDatabase()
	result, err := db.Query("SELECT * from todo")
	defer db.Close()
	return result, err
}

func CreateTodo(id any, title any, completed any) (*sql.Rows, error) {
	db := InitDatabase()
	result, err := db.Query("INSERT INTO todo(Id,Content, Completed) VALUES(?, ?, ?);", id, title, completed)
	defer db.Close()
	return result, err
}

func GetTodo(params string) (*sql.Rows, error) {
	return IsTodoExists(params)
}
func IsTodoExists(id string) (*sql.Rows, error) {
	db := InitDatabase()
	rows, err := db.Query("SELECT * from todo where Id=?", id)
	defer db.Close()
	return rows, err
}
func UpdateTodo(params model.Todo) (int64, error) {
	db := InitDatabase()
	stmt, err := db.Prepare("UPDATE todo SET Content = ?, Completed=? WHERE Id = ?")
	result, err := stmt.Exec(params.Content, params.Completed, params.Id)
	affected, err := result.RowsAffected()

	return affected, err
}

func DeleteTodo(param string) (int64, error) {
	db := InitDatabase()
	stmt, err := db.Prepare("DELETE FROM todo WHERE Id = ?")
	result, err := stmt.Exec(param)
	affected, err := result.RowsAffected()
	return affected, err
}
