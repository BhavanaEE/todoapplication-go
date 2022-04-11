package service

import (
	"encoding/json"
	"io"
	"io/ioutil"
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

func CreateTodo(r io.ReadCloser) {
	responseBody, _ := ioutil.ReadAll(r)
	keyVal := make(map[string]any)
	json.Unmarshal(responseBody, &keyVal)

	id := keyVal["Id"]
	title := keyVal["Content"]
	completed := keyVal["Completed"]
	database.CreateTodo(id, title, completed)
}
