package database

import (
	"database/sql"
	"fmt"
)

func InitDatabase() *sql.DB {
	Db, err := sql.Open("mysql", "root:Bhavs@66@tcp(127.0.0.1:3306)/todo")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Db is connected")
	return Db
}
