package service

import (
	"database/sql"
	"fmt"
	"todoapplication/database"
	"todoapplication/model"
)

type Service struct{}

func (service *Service) GetAllTodos(db *sql.DB) ([]model.Todo, error) {
	repo := &database.Repo{}
	result, err := repo.GetAllTodos(db)
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

func (service *Service) CreateTodo(newTodo model.Todo, db *sql.DB) (int, error) {
	repo := &database.Repo{}
	id := newTodo.Id
	title := newTodo.Content
	completed := newTodo.Completed
	exists, err := repo.IsTodoExists(id, db)
	next := exists.Next()
	fmt.Println(next)
	if err != nil || exists.Next() {
		return 0, err
	}
	todo, err := repo.CreateTodo(id, title, completed, db)
	if todo == 1 && err == nil {
		return 1, err
	}
	return 0, err
}

func (service *Service) GetTodo(id int, db *sql.DB) (model.Todo, error) {
	repo := &database.Repo{}
	var todo model.Todo
	result, err := repo.GetTodo(id, db)
	for result.Next() {
		err := result.Scan(&todo.Id, &todo.Content, &todo.Completed)
		if err != nil {
			panic(err.Error())
		}
	}
	defer result.Close()
	return todo, err
}

func (service *Service) UpdateTodo(id int, todo model.Todo, db *sql.DB) (int, error) {
	repo := &database.Repo{}
	exists, err := repo.IsTodoExists(id, db)
	if err != nil || !exists.Next() {
		return 0, err
	}
	return repo.UpdateTodo(id, todo, db)
}

func (service *Service) DeleteTodo(id int, db *sql.DB) (int, error) {
	repo := &database.Repo{}
	return repo.DeleteTodo(id, db)
}
