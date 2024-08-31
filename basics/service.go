//* contains all our db logic and interactions

package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)


type ToDo struct{
	// ! Capitalize to mark them as exported and accessible fromm outside the package*
	//* Use the json tag to providee metadata* which can be used to convert the entity into JSON when sending data over the network
	Id int `json:"id"`
	Title string `json:"title"`
	Status string `json:"status"`
}

var DB *sql.DB