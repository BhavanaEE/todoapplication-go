package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"todoapplication/controller"
	"todoapplication/database"
	"todoapplication/service"
)

func main() {
	db := database.InitDatabase()

	repo := database.NewRepo(db)
	//service := &service.service{repo}
	newService := service.NewService(repo)

	api := controller.NewController(newService)
	router := gin.Default()
	//newRepo := database.NewRepo(db)
	//newRepo.GetAllTodos()
	router.GET("/todos", api.GetAllTodos)
	router.POST("/todos", api.CreateTodo)
	router.GET("/todos/:Id", api.GetTodo)
	router.PUT("/todos/:Id", api.UpdateTodo)
	router.DELETE("/todos/:Id", api.DeleteTodo)

	router.Run("localhost:8080")
}
