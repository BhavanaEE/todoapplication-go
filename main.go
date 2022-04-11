package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	"todoapplication/rest"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/todos", rest.GetTodos).Methods("GET")

	http.ListenAndServe(":8000", router)
}
