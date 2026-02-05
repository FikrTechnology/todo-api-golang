package service

import (
	"errors"
	"todo-api/model"
	"todo-api/repository"
)

// GetTodos - service untuk mengambil semua todos
func GetTodos() []model.Todo {
	// STEP 4a: Panggil repository untuk mendapatkan semua data
	// STEP 5: Repository return semua data dari storage
	return repository.GetAll()
}

// CreateTodo - service untuk membuat todo baru
func CreateTodo(title string, description string) (model.Todo, error) {
	// STEP 4a: Validasi - pastikan title tidak kosong
	if title == "" {
		return model.Todo{}, errors.New("Title tidak boleh kosong")
	}

	// STEP 4b: Buat object todo baru dengan status false (belum selesai)
	todo := model.Todo{
		Title:       title,
		Done:        false,
		Description: description,
	}

	// STEP 4c: Panggil repository untuk menyimpan todo
	// STEP 5: Repository assign ID dan simpan ke storage
	// STEP 6: Repository return todo dengan ID yang sudah di-assign
	return repository.Create(todo), nil
}

// GetTodoById - service untuk mendapatkan todo berdasarkan id
func GetTodoById(id int) (model.Todo, error) {
	// STEP 4a: Validasi - pastikan id valid (lebih dari 0)
	if id <= 0 {
		return model.Todo{}, errors.New("id tidak valid")
	}

	// STEP 4b: Panggil repository untuk mencari todo dengan id tertentu
	// STEP 5: Repository search di storage
	todo, found := repository.GetById(id)

	// STEP 4c: Validasi - pastikan data ditemukan
	if !found {
		return model.Todo{}, errors.New("Data tidak di temukan!!!")
	}

	// STEP 6: Return todo yang ditemukan
	return todo, nil
}

// UpdateTodo - service untuk mengupdate todo berdasarkan id
func UpdateTodo(id int, title string, description string, done bool) (model.Todo, error) {
	// STEP 4a: Validasi - pastikan id valid (lebih dari 0)
	if id <= 0 {
		return model.Todo{}, errors.New("id tidak valid")
	}

	// STEP 4b: Validasi - pastikan title tidak kosong
	if title == "" {
		return model.Todo{}, errors.New("Title tidak boleh kosong")
	}

	// STEP 4c: Buat object todo dengan data yang diupdate
	updateTodo := model.Todo{
		Title:       title,
		Done:        done,
		Description: description,
	}

	todo, found := repository.Update(id, updateTodo)
	if !found {
		return model.Todo{}, errors.New("todo tidak ditemukan")
	}

	return todo, nil
}
