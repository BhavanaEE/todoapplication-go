package mock

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"todoapplication/model"
	"todoapplication/rest"
	"todoapplication/service"
)

var todo = &model.Todo{Id: 1, Content: "Golang", Completed: false}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock
}

func TestShouldGetAllTodos(t *testing.T) {
	db, mock := NewMock()
	api := &rest.Api{db}
	defer func() { db.Close() }()

	query := "SELECT id, content, completed FROM todo"

	rows := sqlmock.NewRows([]string{"Id", "Content", "Completed"}).
		AddRow(1, "Golang", false).
		AddRow(2, "Java", true)

	mock.ExpectQuery(query).WillReturnRows(rows)

	actual, _ := service.GetAllTodos(api.Db)

	expected := []model.Todo{
		{Id: 1, Content: "Golang", Completed: false},
		{Id: 2, Content: "Java", Completed: true},
	}

	assert.Equal(t, expected, actual, "The two words should be the same.")

}

func TestShouldFailToGetAllTodos(t *testing.T) {
	db, mock := NewMock()
	api := &rest.Api{db}
	defer func() { db.Close() }()

	query := "SELECT id, content, completed FROM todo"

	rows := sqlmock.NewRows([]string{"Id", "Content", "Completed"}).
		AddRow(1, "Golang", false).
		AddRow(2, "Java", true).
		AddRow(3, "Python", true)

	mock.ExpectQuery(query).WillReturnRows(rows)

	actual, _ := service.GetAllTodos(api.Db)

	expected := []model.Todo{
		{Id: 1, Content: "Golang", Completed: false},
		{Id: 2, Content: "Java", Completed: true},
	}

	assert.NotEqualf(t, expected, actual, "The two words should be the same.")

}

func TestShouldGetTodoById(t *testing.T) {
	db, mock := NewMock()
	api := &rest.Api{db}
	defer func() { db.Close() }()

	query := "SELECT id, content, completed FROM todo WHERE id = \\?"

	rows := sqlmock.NewRows([]string{"Id", "Content", "Completed"}).
		AddRow(1, "Golang", false)
	id := 1
	mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)

	actual, _ := service.GetTodo(id, api.Db)

	expected := model.Todo{
		Id: 1, Content: "Golang", Completed: false,
	}

	assert.Equal(t, expected, actual, "The two words should be the same.")
}

func TestShouldReturnEmptyResultForIdDoesNotExist(t *testing.T) {
	db, mock := NewMock()
	api := &rest.Api{db}
	defer func() { db.Close() }()

	query := "SELECT id, content, completed FROM todo WHERE id = \\?"

	rows := sqlmock.NewRows([]string{"Id", "Content", "Completed"})
	id := 2
	mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)

	actual, _ := service.GetTodo(id, api.Db)

	expected := model.Todo{
		Id: 0, Content: "", Completed: false,
	}

	assert.Equal(t, expected, actual, "The two words should be the same.")
}

func TestShouldCreateTodoWhenIdDoesNotExists(t *testing.T) {
	db, mock := NewMock()
	api := &rest.Api{db}
	defer func() { db.Close() }()

	query1 := "SELECT id, content, completed FROM todo WHERE id = \\?"
	rows := sqlmock.NewRows([]string{"Id", "Content", "Completed"})
	id := 1
	mock.ExpectQuery(query1).WithArgs(id).WillReturnRows(rows)

	query2 := "INSERT INTO todo \\(id, content, completed\\) VALUES \\(\\?, \\?, \\?\\)"
	prep := mock.ExpectPrepare(query2)
	prep.ExpectExec().WithArgs(todo.Id, todo.Content, todo.Completed).WillReturnResult(sqlmock.NewResult(0, 1))
	_, err := service.CreateTodo(*todo, api.Db)
	assert.NoError(t, err)

}

func TestShouldNotCreateTodoWhenIdExists(t *testing.T) {
	db, mock := NewMock()
	api := &rest.Api{db}
	defer func() { db.Close() }()

	query1 := "SELECT id, content, completed FROM todo WHERE id = \\?"
	rows := sqlmock.NewRows([]string{"Id", "Content", "Completed"}).
		AddRow(1, "Golang", false)
	id := 1
	mock.ExpectQuery(query1).WithArgs(id).WillReturnRows(rows)

	query2 := "INSERT INTO todo \\(id, content, completed\\) VALUES \\(\\?, \\?, \\?\\)"
	prep := mock.ExpectPrepare(query2)
	prep.ExpectExec().WithArgs(todo.Id, todo.Content, todo.Completed).WillReturnResult(sqlmock.NewResult(0, 1))

	_, err := service.CreateTodo(*todo, api.Db)
	assert.NoError(t, err)

}
