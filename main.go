package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Model
type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	IsDone bool   `json:"isdone"`
}

//Init the todos as a slice of the struct
var todos []Todo

//Get all
func getAllTodos(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

//get one
func getTodo(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Set a params variable and send the request, this gets the params

	//Look for the todo with the id
	for _, selectedTodo := range todos {
		if selectedTodo.ID == params["id"] {
			json.NewEncoder(w).Encode(selectedTodo)
			return
		}
	}
	json.NewEncoder(w).Encode(&Todo{})

}

// Post
func createTodo(w http.ResponseWriter, r *http.Request) {

}

// Update
func updateTodo(w http.ResponseWriter, r *http.Request) {

}

//Delete
func deleteTodo(w http.ResponseWriter, r *http.Request) {

}

func main() {

	//Init the router
	router := mux.NewRouter()

	//Mock data
	todos = append(todos, Todo{ID: "1", Title: "Implement Database", IsDone: false})
	todos = append(todos, Todo{ID: "2", Title: "Use Gin", IsDone: false})
	todos = append(todos, Todo{ID: "3", Title: "Finish all challenges", IsDone: false})

	//Create route handler to stablish the endpints
	router.HandleFunc("/api/todo", getAllTodos).Methods("GET")
	router.HandleFunc("/api/todo/{id}", getTodo).Methods("GET")
	router.HandleFunc("/api/todo", createTodo).Methods("POST")
	router.HandleFunc("/api/todo/{id}", updateTodo).Methods("PUT")
	router.HandleFunc("/api/todo", deleteTodo).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}
