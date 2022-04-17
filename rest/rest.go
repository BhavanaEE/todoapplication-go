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
	todos, err := service.GetAllTodos(a.Db)
	if err != nil {
		c.IndentedJSON(http.StatusServiceUnavailable, gin.H{"message": "Server is not ready to handle the request"})
		return
	}
	c.IndentedJSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var newTodo model.Todo
	err := c.BindJSON(&newTodo)
	if err != nil {
		c.IndentedJSON(http.StatusUnsupportedMediaType, gin.H{"message": "unable to parse body"})
		return
	}
	todo, err := service.CreateTodo(newTodo)
	if todo == 0 || err != nil {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "Todo with provided ID Exists"})
		return
	}
	c.IndentedJSON(http.StatusCreated, newTodo)
}

func GetTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("Id"))
	todo, err := service.GetTodo(id)
	if todo.Id == 0 || err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"message": "Todo Id doesn't exists"})
		return
	}
	c.IndentedJSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	var updateTodo model.Todo
	err := c.BindJSON(&updateTodo)
	params := c.Param("Id")
	if err != nil {
		c.IndentedJSON(http.StatusUnsupportedMediaType, gin.H{"message": "unable to parse body"})
		return
	}
	id, _ := strconv.Atoi(params)
	todo, err := service.UpdateTodo(id, updateTodo)
	if err != nil || todo == 0 {
		c.IndentedJSON(http.StatusConflict, gin.H{"message": "Todo with provided ID doesnt Exists"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "updated"})
}

func DeleteTodo(c *gin.Context) {
	params := c.Param("Id")
	todo, err := service.DeleteTodo(params)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	c.IndentedJSON(http.StatusOK, todo)

}
