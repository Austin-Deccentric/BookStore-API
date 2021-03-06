package model

import (
	"database/sql"
	"fmt"
	"log"
	
)

var connection *sql.DB   // global sql database(Can be accessed by only resources in the model package)

// Connect creates the connection to the database 
func Connect() *sql.DB {
	db, err := sql.Open("mysql", "austin:DPhil28%.?1@tcp(localhost:3306)/learn_api")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err) 
	}
	fmt.Println("Connected to database")
	connection = db
	return db
}