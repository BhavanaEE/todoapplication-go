package database

import (
	"database/sql"
	"todoapplication/model"
)

type Repo struct{}

func (repo *Repo) GetAllTodos(db *sql.DB) (*sql.Rows, error) {
	result, err := db.Query("SELECT id, content, completed FROM todo")
	return result, err
}

func (repo *Repo) CreateTodo(id any, title any, completed any, db *sql.DB) (int64, error) {
	query := "INSERT INTO todo (id, content, completed) VALUES (?, ?, ?)"
	stmt, err := db.Prepare(query)
	result, err := stmt.Exec(id, title, completed)
	affected, err := result.RowsAffected()
	return affected, err
}

func (repo *Repo) GetTodo(params int, db *sql.DB) (*sql.Rows, error) {
	return repo.IsTodoExists(params, db)
}
func (repo *Repo) IsTodoExists(id int, db *sql.DB) (*sql.Rows, error) {
	rows, err := db.Query("SELECT id, content, completed FROM todo WHERE id = ?", id)
	return rows, err
}
func (repo *Repo) UpdateTodo(id int, params model.Todo, db *sql.DB) (int, error) {
	stmt, err := db.Prepare("UPDATE todo SET content = ?, completed = ? WHERE id = ?")
	result, err := stmt.Exec(params.Content, params.Completed, id)
	affected, err := result.RowsAffected()

	return int(affected), err
}

func (repo *Repo) DeleteTodo(id int, db *sql.DB) (int, error) {
	stmt, err := db.Prepare("DELETE FROM todo WHERE id = ?")
	result, err := stmt.Exec(id)
	affected, err := result.RowsAffected()
	return int(affected), err
}
