package rest

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"todoapplication/model"
	"todoapplication/service"
)

type Api struct {
	Db *sql.DB
}

func (a *Api) GetTodos(c *gin.Context) {
	service := &service.Service{}
	todos, err := service.GetAllTodos(a.Db)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, gin.H{"message": "Server is not ready to handle the request"})
		return
	}
	c.IndentedJSON(http.StatusOK, todos)
}

func (a *Api) CreateTodo(c *gin.Context) {
	service := &service.Service{}
	var newTodo model.Todo
	err := c.BindJSON(&newTodo)
	if err != nil {
		c.IndentedJSON(http.StatusUnsupportedMediaType, gin.H{"message": "unable to parse body"})
		return
	}
	todo, err := service.CreateTodo(newTodo, a.Db)
	if todo == 0 || err != nil {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "Todo with provided ID Exists"})
		return
	}
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func (a *Api) GetTodo(c *gin.Context) {
	service := &service.Service{}
	id, _ := strconv.Atoi(c.Param("Id"))
	todo, err := service.GetTodo(id, a.Db)
	if todo.Id == 0 || err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Todo Id doesn't exists"})
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

func (a *Api) UpdateTodo(c *gin.Context) {
	service := &service.Service{}
	var updateTodo model.Todo
	err := c.BindJSON(&updateTodo)
	params := c.Param("Id")
	if err != nil {
		c.IndentedJSON(http.StatusUnsupportedMediaType, gin.H{"message": "unable to parse body"})
		return
	}
	id, _ := strconv.Atoi(params)
	todo, err := service.UpdateTodo(id, updateTodo, a.Db)
	if err != nil || todo == 0 {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "Todo with provided ID doesnt Exists"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "updated"})
}

func (a *Api) DeleteTodo(c *gin.Context) {
	service := &service.Service{}
	params := c.Param("Id")
	id, _ := strconv.Atoi(params)
	todo, err := service.DeleteTodo(id, a.Db)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, todo)

}
