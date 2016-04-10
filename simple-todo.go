package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/bmizerany/pat"
	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Todos []Todo

var mainDB *sql.DB

func main() {

	db, errOpenDB := sql.Open("sqlite3", "todo.db")
	checkErr(errOpenDB)
	mainDB = db

	r := pat.New()
	r.Post("/todos/delete/:id", http.HandlerFunc(deleteByID))
	r.Get("/todos/:id", http.HandlerFunc(getByID))
	r.Post("/todos/:id", http.HandlerFunc(updateByID))
	r.Get("/todos", http.HandlerFunc(getAll))
	r.Post("/todos", http.HandlerFunc(insert))

	http.Handle("/", r)

	log.Print(" Running on 12345")
	err := http.ListenAndServe(":12345", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func getAll(w http.ResponseWriter, r *http.Request) {
	rows, err := mainDB.Query("SELECT * FROM todos")
	checkErr(err)
	var todos Todos
	for rows.Next() {
		var todo Todo
		err = rows.Scan(&todo.ID, &todo.Name)
		checkErr(err)
		todos = append(todos, todo)
	}
	jsonB, _ := json.Marshal(todos)
	fmt.Fprintf(w, "%s", string(jsonB))
}

func getByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get(":id")
	stmt, err := mainDB.Prepare(" SELECT * FROM todos where id = ?")
	checkErr(err)
	rows, _ := stmt.Query(id)
	var todo Todo
	for rows.Next() {
		err = rows.Scan(&todo.ID, &todo.Name)
		checkErr(err)
	}
	jsonB, _ := json.Marshal(todo)
	fmt.Fprintf(w, "%s", string(jsonB))
}

func insert(w http.ResponseWriter, r *http.Request) {

}

func updateByID(w http.ResponseWriter, r *http.Request) {
	//id := r.URL.Query().Get(":id")
}

func deleteByID(w http.ResponseWriter, r *http.Request) {
	//id := r.URL.Query().Get(":id")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
