package service

import (
	"database/sql"
	"todoapplication/database"
	"todoapplication/model"
)

type Service struct {
	getAllTodos  func() ([]model.Todo, error)
	createTodo   func(title any, completed any) (model.Todo, error)
	getTodo      func(id int64) (model.Todo, error)
	isTodoExists func(id int64) (*sql.Rows, error)
	updateTodo   func(id int, params model.Todo) (model.Todo, error)
	deleteTodo   func(id int) error
}

func NewService(repo *database.Repo) *Service {
	return &Service{
		getAllTodos:  repo.GetAllTodos,
		createTodo:   repo.CreateTodo,
		getTodo:      repo.GetTodo,
		isTodoExists: repo.IsTodoExists,
		updateTodo:   repo.UpdateTodo,
		deleteTodo:   repo.DeleteTodo,
	}
}

func (service *Service) GetAllTodos() ([]model.Todo, error) {
	result, err := service.getAllTodos()
	return result, err
}

func (service *Service) CreateTodo(newTodo model.Todo) (model.Todo, error) {
	title := newTodo.Content
	completed := newTodo.Completed
	return service.createTodo(title, completed)
}

func (service *Service) GetTodo(id int64) (model.Todo, error) {
	result, err := service.getTodo(id)
	return result, err
}

func (service *Service) UpdateTodo(id int, todo model.Todo) (model.Todo, error) {
	return service.updateTodo(id, todo)
}

func (service *Service) DeleteTodo(id int) error {
	return service.deleteTodo(id)
}
