package service

import (
	"todoapplication/database"
	"todoapplication/model"
)

func GetAllTodos() []model.Todo {
	result := database.GetAllTodos()
	var todos []model.Todo
	for result.Next() {
		var todo model.Todo
		err := result.Scan(&todo.Id, &todo.Content, &todo.Completed)
		if err != nil {
			panic(err.Error())
		}
		todos = append(todos, todo)
	}
	return todos
}

func CreateTodo(newTodo model.Todo) {
	id := newTodo.Id
	title := newTodo.Content
	completed := newTodo.Completed
	database.CreateTodo(id, title, completed)
}

func GetTodo(params string) model.Todo {
	result := database.GetTodo(params)
	var todo model.Todo
	for result.Next() {
		err := result.Scan(&todo.Id, &todo.Content, &todo.Completed)
		if err != nil {
			panic(err.Error())
		}
	}
	defer result.Close()
	return todo
}

func UpdateTodo(keyvalues model.Todo) {
	database.UpdateTodo(keyvalues)
}

func DeleteTodo(params string) {
	database.DeleteTodo(params)
}
