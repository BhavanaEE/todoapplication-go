package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todoapplication/model"
	"todoapplication/service"
)

type Controller struct {
	getAllTodos func() ([]model.Todo, error)
	createTodo  func(newTodo model.Todo) (model.Todo, error)
	getTodo     func(id int64) (model.Todo, error)
	updateTodo  func(id int, todo model.Todo) (model.Todo, error)
	deleteTodo  func(id int) error
}

func NewController(service *service.Service) Controller {
	return Controller{
		getAllTodos: service.GetAllTodos,
		createTodo:  service.CreateTodo,
		getTodo:     service.GetTodo,
		updateTodo:  service.UpdateTodo,
		deleteTodo:  service.DeleteTodo,
	}
}

func (a *Controller) GetAllTodos(c *gin.Context) {
	todos, err := a.getAllTodos()
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, gin.H{"message": "Server is not ready to handle the request"})
		return
	}
	c.IndentedJSON(http.StatusOK, todos)
}

func (a *Controller) CreateTodo(c *gin.Context) {
	var newTodo model.Todo
	err := c.BindJSON(&newTodo)
	if err != nil {
		c.IndentedJSON(http.StatusUnsupportedMediaType, gin.H{"message": "unable to parse body"})
		return
	}
	todo, err := a.createTodo(newTodo)
	if err != nil {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "Todo with provided ID Exists"})
		return
	}
	c.IndentedJSON(http.StatusCreated, todo)
}

func (a *Controller) GetTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("Id"))
	todo, err := a.getTodo(int64(id))
	if todo.Id == 0 || err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo Id doesn't exists"})
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

func (a *Controller) UpdateTodo(c *gin.Context) {
	var updateTodo model.Todo
	err := c.BindJSON(&updateTodo)
	params := c.Param("Id")
	if err != nil {
		c.IndentedJSON(http.StatusUnsupportedMediaType, gin.H{"message": "unable to parse body"})
		return
	}
	id, _ := strconv.Atoi(params)
	todo, err := a.updateTodo(id, updateTodo)
	if err != nil || todo.Id == 0 {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "Todo with provided ID doesnt Exists"})
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

func (a *Controller) DeleteTodo(c *gin.Context) {
	params := c.Param("Id")
	id, _ := strconv.Atoi(params)
	err := a.deleteTodo(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{})
}
