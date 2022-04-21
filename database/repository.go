package database

import (
	"database/sql"
	"errors"
	"todoapplication/model"
)

type Repo struct {
	Db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db}
}

func (repo *Repo) GetAllTodos() ([]model.Todo, error) {
	result, err := repo.Db.Query("SELECT id, content, completed FROM todo")
	var todos []model.Todo
	for result.Next() {
		var todo model.Todo
		err := result.Scan(&todo.Id, &todo.Content, &todo.Completed)
		if err != nil {
			panic(err.Error())
		}
		todos = append(todos, todo)
	}
	return todos, err
}

func (repo *Repo) CreateTodo(title any, completed any) (model.Todo, error) {
	var todo model.Todo
	query := "INSERT INTO todo (content, completed) VALUES (?, ?)"
	stmt, err := repo.Db.Prepare(query)
	if err != nil {
		return todo, err
	}
	result, _ := stmt.Exec(title, completed)
	if err != nil {
		return todo, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return todo, err
	}
	return repo.GetTodo(id)
}

func (repo *Repo) GetTodo(params int64) (model.Todo, error) {
	rows, err := repo.IsTodoExists(params)
	var todo model.Todo
	for rows.Next() {
		err := rows.Scan(&todo.Id, &todo.Content, &todo.Completed)
		if err != nil {
			panic(err.Error())
		}
	}
	defer rows.Close()
	return todo, err
}
func (repo *Repo) IsTodoExists(id int64) (*sql.Rows, error) {
	rows, err := repo.Db.Query("SELECT id, content, completed FROM todo WHERE id = ?", id)
	return rows, err
}
func (repo *Repo) UpdateTodo(id int, params model.Todo) (model.Todo, error) {
	var todo model.Todo
	exists, err := repo.IsTodoExists(int64(id))
	if err != nil || !exists.Next() {
		return todo, err
	}
	stmt, err := repo.Db.Prepare("UPDATE todo SET content = ?, completed = ? WHERE id = ?")
	_, err = stmt.Exec(params.Content, params.Completed, id)
	if err != nil {
		return todo, err
	}
	todo, err = repo.GetTodo(int64(id))
	if err != nil {
		return todo, err
	}
	return todo, err
}

func (repo *Repo) DeleteTodo(id int) error {
	stmt, err := repo.Db.Prepare("DELETE FROM todo WHERE id = ?")
	if err != nil {
		return err
	}
	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if int(affected) != 1 {
		return errors.New("Could not delete a todo")
	}

	return nil
}
