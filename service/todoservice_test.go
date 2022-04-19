package service

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"todoapplication/model"
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
	defer func() { db.Close() }()

	serviceTest := &Service{}
	query := "SELECT id, content, completed FROM todo"

	rows := sqlmock.NewRows([]string{"Id", "Content", "Completed"}).
		AddRow(1, "Golang", false).
		AddRow(2, "Java", true)

	mock.ExpectQuery(query).WillReturnRows(rows)

	actual, _ := serviceTest.GetAllTodos(db)

	expected := []model.Todo{
		{Id: 1, Content: "Golang", Completed: false},
		{Id: 2, Content: "Java", Completed: true},
	}

	assert.Equal(t, expected, actual, "The two words should be the same.")

}

func TestShouldFailToGetAllTodos(t *testing.T) {
	db, mock := NewMock()
	defer func() { db.Close() }()

	serviceTest := &Service{}
	query := "SELECT id, content, completed FROM todo"

	rows := sqlmock.NewRows([]string{"Id", "Content", "Completed"}).
		AddRow(1, "Golang", false).
		AddRow(2, "Java", true).
		AddRow(3, "Python", true)

	mock.ExpectQuery(query).WillReturnRows(rows)

	actual, _ := serviceTest.GetAllTodos(db)

	expected := []model.Todo{
		{Id: 1, Content: "Golang", Completed: false},
		{Id: 2, Content: "Java", Completed: true},
	}

	assert.NotEqualf(t, expected, actual, "The two words should be the same.")

}

func TestShouldGetTodoById(t *testing.T) {
	db, mock := NewMock()
	defer func() { db.Close() }()

	serviceTest := &Service{}
	query := "SELECT id, content, completed FROM todo WHERE id = \\?"

	rows := sqlmock.NewRows([]string{"Id", "Content", "Completed"}).
		AddRow(1, "Golang", false)
	id := 1
	mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)

	actual, _ := serviceTest.GetTodo(id, db)

	expected := model.Todo{
		Id: 1, Content: "Golang", Completed: false,
	}

	assert.Equal(t, expected, actual, "The two words should be the same.")
}

func TestShouldReturnEmptyResultForIdDoesNotExist(t *testing.T) {
	db, mock := NewMock()
	defer func() { db.Close() }()

	serviceTest := &Service{}
	query := "SELECT id, content, completed FROM todo WHERE id = \\?"

	rows := sqlmock.NewRows([]string{"Id", "Content", "Completed"})
	id := 2
	mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)

	actual, _ := serviceTest.GetTodo(id, db)

	expected := model.Todo{
		Id: 0, Content: "", Completed: false,
	}

	assert.Equal(t, expected, actual, "The two words should be the same.")
}

func TestShouldCreateTodoWhenIdDoesNotExists(t *testing.T) {
	db, mock := NewMock()
	defer func() { db.Close() }()

	serviceTest := &Service{}
	query1 := "SELECT id, content, completed FROM todo WHERE id = \\?"
	rows := sqlmock.NewRows([]string{"Id", "Content", "Completed"})
	id := 1
	mock.ExpectQuery(query1).WithArgs(id).WillReturnRows(rows)

	query2 := "INSERT INTO todo \\(id, content, completed\\) VALUES \\(\\?, \\?, \\?\\)"
	prep := mock.ExpectPrepare(query2)
	prep.ExpectExec().WithArgs(todo.Id, todo.Content, todo.Completed).WillReturnResult(sqlmock.NewResult(0, 1))
	_, err := serviceTest.CreateTodo(*todo, db)
	assert.NoError(t, err)

}

func TestShouldNotCreateTodoWhenIdExists(t *testing.T) {
	db, mock := NewMock()
	defer func() { db.Close() }()

	serviceTest := &Service{}
	query1 := "SELECT id, content, completed FROM todo WHERE id = \\?"
	rows := sqlmock.NewRows([]string{"Id", "Content", "Completed"}).
		AddRow(1, "Golang", false)
	id := 1
	mock.ExpectQuery(query1).WithArgs(id).WillReturnRows(rows)

	query2 := "INSERT INTO todo \\(id, content, completed\\) VALUES \\(\\?, \\?, \\?\\)"
	prep := mock.ExpectPrepare(query2)
	prep.ExpectExec().WithArgs(todo.Id, todo.Content, todo.Completed).WillReturnResult(sqlmock.NewResult(0, 1))

	_, err := serviceTest.CreateTodo(*todo, db)
	assert.NoError(t, err)

}

func TestShouldUpdateTodoIfIdExists(t *testing.T) {
	db, mock := NewMock()
	defer func() { db.Close() }()

	serviceTest := &Service{}
	query1 := "SELECT id, content, completed FROM todo WHERE id = \\?"
	rows := sqlmock.NewRows([]string{"Id", "Content", "Completed"}).
		AddRow(1, "Golang", true)
	id := 1
	mock.ExpectQuery(query1).WithArgs(id).WillReturnRows(rows)

	query2 := "UPDATE todo SET content = \\?, completed = \\? WHERE id = \\?"
	prep := mock.ExpectPrepare(query2)
	prep.ExpectExec().WithArgs(todo.Content, todo.Completed, todo.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	isUpdated, err := serviceTest.UpdateTodo(id, *todo, db)
	assert.NoError(t, err)
	assert.Equal(t, 1, isUpdated)
}

func TestShouldNotUpdateTodo(t *testing.T) {
	db, mock := NewMock()
	defer func() { db.Close() }()

	serviceTest := &Service{}
	query1 := "SELECT id, content, completed FROM todo WHERE id = \\?"
	rows := sqlmock.NewRows([]string{"Id", "Content", "Completed"})
	id := 1
	mock.ExpectQuery(query1).WithArgs(id).WillReturnRows(rows)

	query2 := "UPDATE todo SET content = \\?, completed = \\? WHERE id = \\?"
	prep := mock.ExpectPrepare(query2)
	prep.ExpectExec().WithArgs(todo.Content, todo.Completed, todo.Id).WillReturnResult(sqlmock.NewResult(0, 1))

	isUpdated, err := serviceTest.UpdateTodo(id, *todo, db)
	assert.NoError(t, err)
	assert.Equal(t, 0, isUpdated)
}

func TestShouldDeleteTodoIfIdExists(t *testing.T) {
	db, mock := NewMock()
	defer func() { db.Close() }()

	serviceTest := &Service{}
	query := "DELETE FROM todo WHERE id = \\?"

	id := 1
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(id).WillReturnResult(sqlmock.NewResult(0, 1))

	isDeleted, err := serviceTest.DeleteTodo(id, db)
	assert.NoError(t, err)
	assert.Equal(t, 1, isDeleted)

}

func TestShouldNotEffectTheRowIfIdDoesNotExistsWhileDeleting(t *testing.T) {
	db, mock := NewMock()
	defer func() { db.Close() }()

	serviceTest := &Service{}
	query := "DELETE FROM todo WHERE id = \\?"

	id := 4
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(id).WillReturnResult(sqlmock.NewResult(0, 0))

	isDeleted, err := serviceTest.DeleteTodo(id, db)
	assert.NoError(t, err)
	assert.Equal(t, 0, isDeleted)
}
