package service

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"todoapplication/model"
	"todoapplication/rest"
)

func TestShouldGetTodos(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	api := &rest.Api{db}

	rows := sqlmock.NewRows([]string{"Id", "Content", "Completed"}).
		AddRow(1, "Golang", true)
	//mock.ExpectQuery("^SELECT * FROM todo where Id=?").WillReturnRows(rows)
	mock.ExpectQuery("^SELECT (.+) FROM todo$").WillReturnRows(rows)

	expected := []model.Todo{
		{Id: 1, Content: "Golang", Completed: false},
		{Id: 2, Content: "Java", Completed: true},
	}

	//actual, err := GetTodo(1, api.Db)
	actual, err := GetAllTodos(api.Db)

	assert.Equal(t, expected, actual, "The two words should be the same.")

	//w := httptest.NewRecorder()
	//c, _ := gin.CreateTestContext(w)
	//c.Request, _ = http.NewRequest("GET", "http://localhost:8080/todos/1", nil)

	//api.GetTodo(c)

	//if w.Code != 200 {
	//	t.Fatalf("expected status code to be 200, but got: %d", w.Code)
	//}

	//if err := mock.ExpectationsWereMet(); err != nil {
	//	t.Errorf("there were unfulfilled expectations: %s", err)
	//}
}
