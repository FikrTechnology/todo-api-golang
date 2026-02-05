package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"todo-api/service"
)

// GetTodos - handler untuk GET /todos
func GetTodos(w http.ResponseWriter, r *http.Request) {
	// STEP 3a: Set response header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// STEP 3b: Panggil service untuk mendapatkan semua todos
	// STEP 4: Service memanggil repository untuk fetch data
	// STEP 5: Repository mengembalikan data dari storage
	// STEP 6: Encode response sebagai JSON dan kirim ke client
	json.NewEncoder(w).Encode(service.GetTodos())
}

// CreateTodo - handler untuk POST /todos
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	// STEP 3a: Siapkan struct untuk menerima request body
	var body struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}

	// STEP 3b: Decode JSON dari request body ke struct
	json.NewDecoder(r.Body).Decode(&body)

	// STEP 3c: Panggil service untuk membuat todo baru
	// STEP 4: Service melakukan validasi input
	// STEP 5: Service membuat object todo dan panggil repository
	// STEP 6: Repository menyimpan data dan return object
	todo, err := service.CreateTodo(body.Title, body.Description)

	// STEP 7: Cek error dari service
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// STEP 8: Set status code HTTP 201 (Created)
	w.WriteHeader(http.StatusCreated)

	// STEP 9: Encode object todo sebagai JSON response
	json.NewEncoder(w).Encode(todo)
}

// GetTodoById - handler untuk GET /todos/{id}
func GetTodoById(w http.ResponseWriter, r *http.Request) {
	// STEP 3a: Extract id dari URL path
	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")

	// STEP 3b: Convert string id ke integer
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "id harus berupa angka", http.StatusBadRequest)
		return
	}

	// STEP 3c: Panggil service untuk mencari todo berdasarkan id
	// STEP 4: Service melakukan validasi id
	// STEP 5: Service panggil repository untuk mencari data
	// STEP 6: Repository search di storage dan return hasil
	todo, err := service.GetTodoById(id)

	// STEP 7: Cek error dari service
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
	}

	// STEP 8: Set response header sebagai JSON
	w.Header().Set("Content-Type", "application/json")

	// STEP 9: Encode object todo sebagai JSON response
	json.NewEncoder(w).Encode(todo)
}
