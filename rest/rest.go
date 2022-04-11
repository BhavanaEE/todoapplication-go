package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"todoapplication/service"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := service.GetAllTodos()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	service.CreateTodo(r.Body)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "New todo was created")
}
