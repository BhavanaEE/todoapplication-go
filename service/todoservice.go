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
