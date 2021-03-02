package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/joshuaherrera/go-todo/server/models"
)

// Res is a struct for simple responses
type Res struct {
	Message string `json:"message"`
}

func (app *application) home(w http.ResponseWriter, r *http.Request)  {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	// data := []byte(`{"mesage":"hello world!"}`)
	// w.Write(data)

	data := Res{"hello world!"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(data)
}

func (app *application) getTodos(w http.ResponseWriter, r *http.Request) {
	t, err := app.todos.All()
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(t)
}

func (app *application) getTodo(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	todoID, err := strconv.Atoi(vars["id"])

	if err != nil || todoID < 1 {
		app.notFound(w)
		return
	}

	t, err := app.todos.Get(todoID)
	if err != nil {
		app.serverError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(t)

}

func (app *application) insertTodo(w http.ResponseWriter, r *http.Request)  {
	// read body
	todo := models.Todo{}

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	t, err := app.todos.Insert(todo.Description)
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(t)
}

func (app *application) updateTodo(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	todoID, err := strconv.Atoi(vars["id"])
	if err != nil || todoID < 1 {
		app.notFound(w)
		return
	}
	todo := models.Todo{}

	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	err = app.todos.Update(todoID, todo.Description)
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := Res{"Todo was updated"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(data)

}

func (app *application) deleteTodo(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	todoID, err := strconv.Atoi(vars["id"])
	if err != nil || todoID < 1 {
		app.notFound(w)
		return
	}

	err = app.todos.Delete(todoID)
	if err != nil {
		app.serverError(w, err)
		return
	}

	data := Res{"Todo was deleted"}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(data)

}