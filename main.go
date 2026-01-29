package main

import (
	"net/http"
	"todo-api/handler"
)

func main() {
	// STEP 1: Setup route untuk endpoint /todos
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		// STEP 2a: Handle GET request - ambil semua todos
		if r.Method == http.MethodGet {
			handler.GetTodos(w, r)
		}

		// STEP 2b: Handle POST request - buat todo baru
		if r.Method == http.MethodPost {
			handler.CreateTodo(w, r)
		}
	})

	// STEP 1: Setup route untuk endpoint /todos/{id}
	http.HandleFunc("/todos/", func(w http.ResponseWriter, r *http.Request) {
		// STEP 2: Handle GET request - ambil todo berdasarkan id
		if r.Method == http.MethodGet {
			handler.GetTodoById(w, r)
		}
	})

	// STEP 3: Start HTTP server di port 8000
	http.ListenAndServe(":8000", nil)
}
