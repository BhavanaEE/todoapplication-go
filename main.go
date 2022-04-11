package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"todoapplication/rest"
)

func main() {
	router := gin.Default()
	router.GET("/todos", rest.GetTodos)
	router.POST("/todos", rest.CreateTodo)
	router.GET("/todos/:Id", rest.GetTodo)
	router.PUT("/todos/:Id", rest.UpdateTodo)
	router.DELETE("/todos/:Id", rest.DeleteTodo)
	router.Run("localhost:8080")
}
