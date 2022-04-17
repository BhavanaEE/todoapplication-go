package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"todoapplication/database"
	"todoapplication/rest"
)

func main() {
	db := database.InitDatabase()
	api := &rest.Api{db}
	router := gin.Default()
	router.GET("/todos", api.GetTodos)
	router.POST("/todos", api.CreateTodo)
	router.GET("/todos/:Id", api.GetTodo)
	router.PUT("/todos/:Id", api.UpdateTodo)
	router.DELETE("/todos/:Id", api.DeleteTodo)

	router.Run("localhost:8080")
}
