package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {

	router := mux.NewRouter()
	// Middleware
	router.Use(mux.CORSMethodMiddleware(router))
	// r.Use(limit)
	router.Use(allowCors)


	router.HandleFunc("/", app.home)
	router.HandleFunc("/todos", app.getTodos).Methods(http.MethodGet)
	router.HandleFunc("/todos", app.insertTodo).Methods(http.MethodPost, http.MethodOptions)
	router.HandleFunc(("/todos/{id:[0-9]+}"), app.getTodo).Methods(http.MethodGet)
	router.HandleFunc(("/todos/{id:[0-9]+}"), app.updateTodo).Methods(http.MethodPut, http.MethodOptions)
	router.HandleFunc(("/todos/{id:[0-9]+}"), app.deleteTodo).Methods(http.MethodDelete, http.MethodOptions)



	return router
}

func allowCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Access-Control-Allow-Headers, Authorization, X-Requested-With")
			return
		}
		next.ServeHTTP(w, r)
	})
}