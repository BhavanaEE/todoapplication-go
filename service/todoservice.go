package service

import (
	"database/sql"
	"todoapplication/database"
	"todoapplication/model"
)

func GetAllTodos(db *sql.DB) ([]model.Todo, error) {
	result, err := database.GetAllTodos(db)
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

func CreateTodo(newTodo model.Todo) (int, error) {
	id := newTodo.Id
	title := newTodo.Content
	completed := newTodo.Completed
	exists, err := database.IsTodoExists(id)
	if err != nil || !exists.Next() {
		return 0, err
	}
	todo, err := database.CreateTodo(id, title, completed)
	if todo.Next() || err != nil {
		return 1, err
	}
	return 0, err
}

func GetTodo(id int) (model.Todo, error) {
	var todo model.Todo
	result, err := database.GetTodo(id)
	for result.Next() {
		err := result.Scan(&todo.Id, &todo.Content, &todo.Completed)
		if err != nil {
			panic(err.Error())
		}
	}
	defer result.Close()
	return todo, err
}

func UpdateTodo(id int, keyvalues model.Todo) (int64, error) {
	exists, err := database.IsTodoExists(id)
	if err != nil || !exists.Next() {
		return 0, err
	}
	return database.UpdateTodo(keyvalues)
}

func DeleteTodo(params string) (int64, error) {
	return database.DeleteTodo(params)
}
