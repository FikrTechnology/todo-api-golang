package repository

import "todo-api/model"

// STEP 5: Data storage (in-memory)
var todos = []model.Todo{}

// GetAll - repository function untuk mendapatkan semua todos
func GetAll() []model.Todo {
	// STEP 5a: Return semua data dari storage
	// Data ini akan di-return ke service, lalu ke handler, lalu di-encode ke JSON
	return todos
}

// Create - repository function untuk menyimpan todo baru
func Create(todo model.Todo) model.Todo {
	// STEP 5a: Assign ID berdasarkan jumlah data yang sudah ada
	todo.ID = len(todos) + 1

	// STEP 5b: Tambahkan todo ke storage (in-memory slice)
	todos = append(todos, todo)

	// STEP 5c: Return todo yang sudah di-simpan dengan ID
	return todo
}

// GetById - repository function untuk mencari todo berdasarkan id
func GetById(id int) (model.Todo, bool) {
	// STEP 5a: Loop melalui semua data di storage
	for _, todo := range todos {
		// STEP 5b: Cek apakah ID cocok
		if todo.ID == id {
			// STEP 5c: Return todo dan flag true (data ditemukan)
			return todo, true
		}
	}

	// STEP 5d: Return empty todo dan flag false (data tidak ditemukan)
	return model.Todo{}, false
}

// Update data - repository function untuk mengupdate todo berdasarkan id
func Update(id int, updateTodo model.Todo) (model.Todo, bool) {
	// STEP 5a: Loop melalui semua data di storage
	for i, todo := range todos {
		// STEP 5b: Cek apakah ID cocok
		if todo.ID == id {
			// STEP 5c: Update data di storage
			updateTodo.ID = id
			todos[i] = updateTodo

			// STEP 5d: Return todo yang sudah di-update dan flag true (data ditemukan)
			return updateTodo, true
		}
	}

	// STEP 5e: Return empty todo dan flag false (data tidak ditemukan)
	return model.Todo{}, false
}
