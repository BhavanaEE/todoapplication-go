package rest

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
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

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err.Error())
	}
	keyVal := make(map[string]any)
	json.Unmarshal(body, &keyVal)
	service.UpdateTodo(keyVal)
	fmt.Fprintf(w, "Todo with ID = %v was updated", keyVal["Id"])
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	service.DeleteTodo(params["Id"])

	fmt.Fprintf(w, "Todo with ID = %s was deleted", params["Id"])
}
