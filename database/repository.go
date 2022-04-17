package database

import (
	"database/sql"
	"todoapplication/model"
)

func GetAllTodos(db *sql.DB) (*sql.Rows, error) {
	result, err := db.Query("SELECT id, content, completed FROM todo")
	defer db.Close()
	return result, err
}

func CreateTodo(id any, title any, completed any, db *sql.DB) (int64, error) {
	query := "INSERT INTO todo (id, content, completed) VALUES (?, ?, ?)"
	stmt, err := db.Prepare(query)
	result, err := stmt.Exec(id, title, completed)
	affected, err := result.RowsAffected()
	defer db.Close()
	return affected, err
}

func GetTodo(params int, db *sql.DB) (*sql.Rows, error) {
	return IsTodoExists(params, db)
}
func IsTodoExists(id int, db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("SELECT id, content, completed FROM todo WHERE id = ?", id)
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
