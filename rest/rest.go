package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todoapplication/model"
	"todoapplication/service"
)

func GetTodos(c *gin.Context) {
	todos := service.GetAllTodos()
	c.IndentedJSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var newTodo model.Todo
	err := c.BindJSON(&newTodo)
	if err != nil {
		return
	}
	service.CreateTodo(newTodo)
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func GetTodo(c *gin.Context) {
	id := c.Param("Id")
	todo := service.GetTodo(id)
	c.IndentedJSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	var updateTodo model.Todo
	err := c.BindJSON(&updateTodo)
	if err != nil {
		return
	}
	if err != nil {
		panic(err.Error())
	}
	service.UpdateTodo(updateTodo)
}

func DeleteTodo(c *gin.Context) {
	params := c.Param("Id")
	service.DeleteTodo(params)
}
