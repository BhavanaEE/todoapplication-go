package database

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

	repo := Repo{db}
	query := "SELECT id, content, completed FROM todo"

	rows := sqlmock.NewRows([]string{"Id", "Content", "Completed"}).
		AddRow(1, "Golang", false).
		AddRow(2, "Java", true)

	mock.ExpectQuery(query).WillReturnRows(rows)

	actual, _ := repo.GetAllTodos()

	expected := []model.Todo{
		{Id: 1, Content: "Golang", Completed: false},
		{Id: 2, Content: "Java", Completed: true},
	}

	assert.Equal(t, expected, actual, "The two words should be the same.")

}

func TestShouldFailToGetAllTodos(t *testing.T) {
	db, mock := NewMock()
	defer func() { db.Close() }()

	repo := Repo{db}
	query := "SELECT id, content, completed FROM todo"

	rows := sqlmock.NewRows([]string{"Id", "Content", "Completed"}).
		AddRow(1, "Golang", false).
		AddRow(2, "Java", true).
		AddRow(3, "Python", true)

	mock.ExpectQuery(query).WillReturnRows(rows)

	actual, _ := repo.GetAllTodos()

	expected := []model.Todo{
		{Id: 1, Content: "Golang", Completed: false},
		{Id: 2, Content: "Java", Completed: true},
	}

	assert.NotEqualf(t, expected, actual, "The two words should be the same.")

}

func TestShouldGetTodoById(t *testing.T) {
	db, mock := NewMock()
	defer func() { db.Close() }()

	repo := Repo{db}
	query := "SELECT id, content, completed FROM todo WHERE id = \\?"

	rows := sqlmock.NewRows([]string{"Id", "Content", "Completed"}).
		AddRow(1, "Golang", false)
	id := 1
	mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)

	actual, _ := repo.GetTodo(int64(id))

	expected := model.Todo{
		Id: 1, Content: "Golang", Completed: false,
	}

	assert.Equal(t, expected, actual, "The two words should be the same.")
}

func TestShouldReturnEmptyResultForIdDoesNotExist(t *testing.T) {
	db, mock := NewMock()
	defer func() { db.Close() }()

	repo := Repo{db}
	query := "SELECT id, content, completed FROM todo WHERE id = \\?"

	rows := sqlmock.NewRows([]string{"Id", "Content", "Completed"})
	id := 2
	mock.ExpectQuery(query).WithArgs(id).WillReturnRows(rows)

	actual, _ := repo.GetTodo(int64(id))

	expected := model.Todo{
		Id: 0, Content: "", Completed: false,
	}

	assert.Equal(t, expected, actual, "The two words should be the same.")
}

func TestShouldCreateTodoWhenIdDoesNotExists(t *testing.T) {
	db, mock := NewMock()
	defer func() { db.Close() }()

	repo := Repo{db}
	insertQuery := "INSERT INTO todo \\(content, completed\\) VALUES \\(\\?, \\?\\)"
	prep := mock.ExpectPrepare(insertQuery)
	prep.ExpectExec().WithArgs(todo.Content, todo.Completed).WillReturnResult(sqlmock.NewResult(2, 1))

	getQuery := "SELECT id, content, completed FROM todo WHERE id = \\?"
	rows := sqlmock.NewRows([]string{"Id", "Content", "Completed"})
	id := 2
	mock.ExpectQuery(getQuery).WithArgs(id).WillReturnRows(rows)

	_, err := repo.CreateTodo(todo.Content, todo.Completed)
	assert.NoError(t, err)

}

//func TestShouldNotCreateTodoWhenIdExists(t *testing.T) {
//	db, mock := NewMock()
//	defer func() { db.Close() }()
//
//	repo := Repo{db}
//	insertQuery := "INSERT INTO todo \\(content, completed\\) VALUES \\(\\?, \\?\\)"
//	prep := mock.ExpectPrepare(insertQuery)
//	prep.ExpectExec().WithArgs(todo.Content, todo.Completed).WillReturnResult(sqlmock.NewResult(2, 1))
//
//	getQuery := "SELECT id, content, completed FROM todo WHERE id = \\?"
//	rows := sqlmock.NewRows([]string{"Id", "Content", "Completed"})
//	id := 2
//	mock.ExpectQuery(getQuery).WithArgs(id).WillReturnRows(rows)
//
//	_, err := repo.CreateTodo(todo.Content, todo.Completed)
//	assert.NoError(t, err)
//
//	assert.NoError(t, err)
//
//}

func TestShouldUpdateTodoIfIdExists(t *testing.T) {
	db, mock := NewMock()
	defer func() { db.Close() }()

	repo := Repo{db}
	isTodoExistsQuery := "SELECT id, content, completed FROM todo WHERE id = \\?"
	rows := sqlmock.NewRows([]string{"Id", "Content", "Completed"}).
		AddRow(1, "Golang", true)
	id := 1
	mock.ExpectQuery(isTodoExistsQuery).WithArgs(id).WillReturnRows(rows)

	updateQuery := "UPDATE todo SET content = \\?, completed = \\? WHERE id = \\?"
	prep := mock.ExpectPrepare(updateQuery)
	prep.ExpectExec().WithArgs(todo.Content, todo.Completed, todo.Id).WillReturnResult(sqlmock.NewResult(1, 1))

	getQuery := "SELECT id, content, completed FROM todo WHERE id = \\?"
	row := sqlmock.NewRows([]string{"Id", "Content", "Completed"}).
		AddRow(1, "Golang", false)
	mock.ExpectQuery(getQuery).WithArgs(id).WillReturnRows(row)

	updateTodo, err := repo.UpdateTodo(id, *todo)

	assert.NoError(t, err)
	assert.Equal(t, *todo, updateTodo)
}

func TestShouldNotUpdateTodoIfIdDoesNotExists(t *testing.T) {
	db, mock := NewMock()
	defer func() { db.Close() }()

	repo := Repo{db}
	isTodoExistsQuery := "SELECT id, content, completed FROM todo WHERE id = \\?"
	rows := sqlmock.NewRows([]string{"Id", "Content", "Completed"})
	id := 1
	mock.ExpectQuery(isTodoExistsQuery).WithArgs(id).WillReturnRows(rows)

	updateTodo, err := repo.UpdateTodo(id, *todo)

	assert.NoError(t, err)
	assert.Equal(t, model.Todo{}, updateTodo)
}

func TestShouldDeleteTodoIfIdExists(t *testing.T) {
	db, mock := NewMock()
	defer func() { db.Close() }()

	repo := Repo{db}
	query := "DELETE FROM todo WHERE id = \\?"

	id := 1
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(id).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.DeleteTodo(id)
	assert.NoError(t, err)

}

func TestShouldNotEffectTheRowIfIdDoesNotExistsWhileDeleting(t *testing.T) {
	db, mock := NewMock()
	defer func() { db.Close() }()

	repo := Repo{db}
	query := "DELETE FROM todo WHERE id = \\?"

	id := 4
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(id).WillReturnResult(sqlmock.NewResult(0, 0))

	err := repo.DeleteTodo(id)
	assert.Error(t, err)
}
