package service

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"todoapplication/model"
)

var todo = model.Todo{Id: 1, Content: "Golang", Completed: false}

func TestShouldGetAllTodos(t *testing.T) {
	serviceTest := &Service{
		getAllTodos: func() ([]model.Todo, error) {
			actual := []model.Todo{
				{Id: 1, Content: "Golang", Completed: false},
				{Id: 2, Content: "Java", Completed: true},
			}
			return actual, nil
		},
	}

	actual, _ := serviceTest.GetAllTodos()

	expected := []model.Todo{
		{Id: 1, Content: "Golang", Completed: false},
		{Id: 2, Content: "Java", Completed: true},
	}

	assert.Equal(t, expected, actual, "The two words should be the same.")

}

func TestShouldFailToGetAllTodos(t *testing.T) {
	serviceTest := &Service{
		getAllTodos: func() ([]model.Todo, error) {
			actual := []model.Todo{
				{Id: 1, Content: "Golang", Completed: false},
				{Id: 2, Content: "Java", Completed: true},
			}
			return actual, nil
		},
	}

	actual, _ := serviceTest.GetAllTodos()

	expected := []model.Todo{
		{Id: 1, Content: "Golang", Completed: false},
	}

	assert.NotEqual(t, expected, actual, "The two words should be the same.")

}

func TestShouldGetTodoById(t *testing.T) {
	serviceTest := &Service{
		getTodo: func(id int64) (model.Todo, error) {
			actual := model.Todo{
				Id: 1, Content: "Golang", Completed: false}
			return actual, nil
		},
	}

	actual, _ := serviceTest.GetTodo(1)

	expected := model.Todo{
		Id: 1, Content: "Golang", Completed: false,
	}

	assert.Equal(t, expected, actual, "The two words should be the same.")

}

func TestShouldReturnEmptyResultForIdDoesNotExist(t *testing.T) {
	serviceTest := &Service{
		getTodo: func(id int64) (model.Todo, error) {
			actual := model.Todo{}
			return actual, nil
		},
	}

	actual, _ := serviceTest.GetTodo(10)

	expected := model.Todo{}

	assert.Equal(t, expected, actual, "The two words should be the same.")

}

func TestShouldCreateTodoWhenIdDoesNotExists(t *testing.T) {
	serviceTest := &Service{
		createTodo: func(title any, completed any) (model.Todo, error) {
			actual := model.Todo{
				Id: 1, Content: "Golang", Completed: false}
			return actual, nil
		},
	}

	actual, _ := serviceTest.createTodo("Golang", false)

	expected := model.Todo{
		Id: 1, Content: "Golang", Completed: false,
	}

	assert.Equal(t, expected, actual, "The two words should be the same.")

}

func TestShouldUpdateTodoIfIdExists(t *testing.T) {
	serviceTest := &Service{
		updateTodo: func(id int, params model.Todo) (model.Todo, error) {
			actual := model.Todo{
				Id: 1, Content: "Golang", Completed: true}
			return actual, nil
		},
	}

	actual, _ := serviceTest.UpdateTodo(1, todo)

	expected := model.Todo{
		Id: 1, Content: "Golang", Completed: true,
	}

	assert.Equal(t, expected, actual, "The two words should be the same.")
}

func TestShouldDeleteTodoIfIdExists(t *testing.T) {
	serviceTest := &Service{
		deleteTodo: func(id int) error {
			return nil
		},
	}

	actual := serviceTest.deleteTodo(1)

	expected := error(nil)

	assert.Equal(t, expected, actual, "The two words should be the same.")

}

func TestShouldNotEffectTheRowIfIdDoesNotExistsWhileDeleting(t *testing.T) {
	serviceTest := &Service{
		deleteTodo: func(id int) error {
			return errors.New("Id doesn't exists")
		},
	}

	actual := serviceTest.deleteTodo(10)

	expected := errors.New("Id doesn't exists")

	assert.Equal(t, expected, actual, "The two words should be the same.")

}
