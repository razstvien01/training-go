//* contains all our db logic and interactions

package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//* Type or interface of ToDo
type ToDo struct{
	// ! Capitalize to mark them as exported and accessible fromm outside the package*
	//* Use the json tag to providee metadata* which can be used to convert the entity into JSON when sending data over the network
	Id int `json:"id"`
	Title string `json:"title"`
	Status string `json:"status"`
}

var DB *sql.DB

func InitDB(){
	var err error
	DB, err = sql.Open("sqlite3", "./awesome.db")
	
	if err != nil {
		log.Fatal(err)
	}
	
	//* Creating todos table
	_, err = DB.Exec(`
	CREATE TABLE IF NOT EXISTS todos (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		status TEXT
	);`)
	
	if err != nil{
		log.Fatal(err)
	}
}

func CreateToDo(title string, status string) (int64, error){
	result, err := DB.Exec(
		"INSERT INTO todos (title, status) VALUES (?, ?)",
		title, status,
	)
	
	if err != nil{
		return 0, err
	}
	
	id, err := result.LastInsertId()
	
	if err != nil {
		return 0, err
	}
	
	return id, nil
}

func DeleteToDo(id int64) error{
	_, err := DB.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}

func ReadToDoList() []ToDo {
	rows, _ := DB.Query("SELECT id, title, status FROM todos")
	
	defer rows.Close()
	
	todos := make([]ToDo, 0)
	
	for rows.Next(){
		var todo ToDo
		rows.Scan(&todo.Id, &todo.Title, &todo.Status)
		todos = append(todos, todo)
	}
	
	return todos
}