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
	router.GET("/todos/:Id", api.GetTodo)
	router.Run("localhost:8080")
}
