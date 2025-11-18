package main

import (
	"encoding/json"
	"log"
	"net/http"
)


// first i need to create a "library" storing the values that each todo will have, itll need its own data structure
type TodoItem struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

// then i need hese todos to be added to a list , so each element can be taken and displayed as an individual to do
// declare the datatype of the list aswell

var AllTodos = []TodoItem{

	{ID: 1, Title: "Create A ToDo list", Description: "Go through and create a functioning to do list"},
}

var nextID = 2

func main() {

	http.HandleFunc("/", ToDoListHandler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)

	}

}

func ToDoListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// first we need to manage a get HTTP Method, this takes data from api
	if r.Method == "GET" {
		// we are going to send JSON
		w.Header().Set("Content-type", "application/json")

		//here we write data into the big array of all todos that was created, that points to specific type
		json.NewEncoder(w).Encode(AllTodos)

		return
	}

	// now we need to apply a post http request, this sends requests to the api based on user inputs

	if r.Method == "POST" {

		// we create a new item that will be added to the todo list
		var newItem TodoItem

		// now we need to read the incoming request to add the data to our variable, this way it can be added to the todolist

		err := json.NewDecoder(r.Body).Decode(&newItem)

		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		//now we give the todo a new ID based on the one we set previously, then we incrament the ID

		newItem.ID = nextID
		nextID++

		// add the newItem to the big todo list
		AllTodos = append(AllTodos, newItem)

		// A line to test if values are being added to the list
		log.Printf("new todo: %+v\n", newItem)

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-type", "application/json")

		json.NewEncoder(w).Encode(newItem)

	}

}


