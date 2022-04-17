package service

import (
	"database/sql"
	"fmt"
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

func CreateTodo(newTodo model.Todo, db *sql.DB) (int, error) {
	id := newTodo.Id
	title := newTodo.Content
	completed := newTodo.Completed
	exists, err := database.IsTodoExists(id, db)
	next := exists.Next()
	fmt.Println(next)
	if err != nil || exists.Next() {
		return 0, err
	}
	todo, err := database.CreateTodo(id, title, completed, db)
	if todo == 1 && err == nil {
		return 1, err
	}
	return 0, err
}

func GetTodo(id int, db *sql.DB) (model.Todo, error) {
	var todo model.Todo
	result, err := database.GetTodo(id, db)
	for result.Next() {
		err := result.Scan(&todo.Id, &todo.Content, &todo.Completed)
		if err != nil {
			panic(err.Error())
		}
	}
	defer result.Close()
	return todo, err
}

func UpdateTodo(id int, todo model.Todo, db *sql.DB) (int, error) {
	exists, err := database.IsTodoExists(id, db)
	if err != nil || !exists.Next() {
		return 0, err
	}
	return database.UpdateTodo(todo, db)
}

func DeleteTodo(id int, db *sql.DB) (int, error) {
	return database.DeleteTodo(id, db)
}
