package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
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

func GetTodo(w http.ResponseWriter, r *http.Request) {
	todo := service.GetTodo(mux.Vars(r))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
