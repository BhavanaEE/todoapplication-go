package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todoapplication/model"
)

var todos = []model.Todo{
	{Id: 1, Content: "Golang", Completed: false},
	{Id: 2, Content: "Java", Completed: true},
}

var todo = model.Todo{Id: 1, Content: "Golang", Completed: false}

func TestShouldGetAllTodos(t *testing.T) {
	c := &Controller{getAllTodos: func() ([]model.Todo, error) {
		return todos, nil
	}}
	w := httptest.NewRecorder()
	r := gin.New()

	r.GET("/todos", func(context *gin.Context) {
		c.GetAllTodos(context)
	})

	req, _ := http.NewRequest("GET", "/todos", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestShouldFailGetAllTodos(t *testing.T) {
	c := &Controller{getAllTodos: func() ([]model.Todo, error) {
		return []model.Todo{}, errors.New("failed")
	}}
	w := httptest.NewRecorder()
	r := gin.New()

	r.GET("/todos", func(context *gin.Context) {
		c.GetAllTodos(context)
	})

	req, _ := http.NewRequest("GET", "/todos", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusServiceUnavailable, w.Code)
	p, _ := ioutil.ReadAll(w.Body)
	assert.Equal(t, "{\n    \"message\": \"Server is not ready to handle the request\"\n}", string(p))
}

func TestShouldGetTodoById(t *testing.T) {
	c := &Controller{getTodo: func(id int64) (model.Todo, error) {
		return model.Todo{int(id), "Java", true}, nil
	}}
	w := httptest.NewRecorder()
	r := gin.New()

	r.GET("/todos/:Id", func(context *gin.Context) {
		c.GetTodo(context)
	})

	req, _ := http.NewRequest("GET", "/todos/1", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestShouldFailToGetTodoById(t *testing.T) {
	c := &Controller{getTodo: func(int642 int64) (model.Todo, error) {
		return model.Todo{}, errors.New("failed")
	}}
	w := httptest.NewRecorder()
	r := gin.New()

	r.GET("/todos/:Id", func(context *gin.Context) {
		c.GetTodo(context)
	})

	req, _ := http.NewRequest("GET", "/todos/1", nil)

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestShouldCreateTodo(t *testing.T) {
	c := &Controller{createTodo: func(newTodo model.Todo) (model.Todo, error) {
		return newTodo, nil
	}}
	w := httptest.NewRecorder()
	r := gin.New()

	r.POST("/todos", func(context *gin.Context) {
		c.CreateTodo(context)
	})

	todoPayload := `{
	   "Id": 11,
	   "Content": "XYZ",
	   "Completed": false
	}`
	req, _ := http.NewRequest("POST", "/todos", strings.NewReader(todoPayload))

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

}

func TestShouldUpdateTodo(t *testing.T) {
	c := &Controller{updateTodo: func(id int, todo model.Todo) (model.Todo, error) {
		return todo, nil
	}}
	w := httptest.NewRecorder()
	r := gin.New()

	r.PUT("/todos/:Id", func(context *gin.Context) {
		c.UpdateTodo(context)
	})

	todoPayload := `{
	   "Id": 11,
	   "Content": "XYZ",
	   "Completed": false
	}`
	req, _ := http.NewRequest("PUT", "/todos/1", strings.NewReader(todoPayload))

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

}

func TestShouldFailToUpdateTodoIfIdDoesNorExists(t *testing.T) {
	c := &Controller{updateTodo: func(id int, todo model.Todo) (model.Todo, error) {
		return model.Todo{}, nil
	}}
	w := httptest.NewRecorder()
	r := gin.New()

	r.PUT("/todos/:Id", func(context *gin.Context) {
		c.UpdateTodo(context)
	})

	todoPayload := `{
	   "Id": 11,
	   "Content": "XYZ",
	   "Completed": false
	}`
	req, _ := http.NewRequest("PUT", "/todos/1", strings.NewReader(todoPayload))

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusConflict, w.Code)

}

func TestShouldDeleteTodoIfIdExists(t *testing.T) {
	c := &Controller{deleteTodo: func(id int) error {
		return nil
	}}
	w := httptest.NewRecorder()
	r := gin.New()

	r.DELETE("/todos/:Id", func(context *gin.Context) {
		c.DeleteTodo(context)
	})
	req, _ := http.NewRequest("DELETE", "/todos/1", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
}

func TestShouldDeleteTodoIfIdDoesNotExists(t *testing.T) {
	c := &Controller{deleteTodo: func(id int) error {
		return errors.New("Not found")
	}}
	w := httptest.NewRecorder()
	r := gin.New()

	r.DELETE("/todos/:Id", func(context *gin.Context) {
		c.DeleteTodo(context)
	})
	req, _ := http.NewRequest("DELETE", "/todos/1", nil)
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusNotFound, w.Code)
}
