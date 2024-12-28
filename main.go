package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

// NOTE: 簡略化のためメモリ上に保存する
var todos = NewTodoList()
var mu sync.Mutex

func main() {
	port := os.Getenv("APP_PORT")
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: nil,
	}

	http.HandleFunc("GET /todos", getTodos)
	http.HandleFunc("POST /todos", addTodo)
	http.HandleFunc("GET /todos/{id}", getTodo)
	http.HandleFunc("PUT /todos/{id}", updateTodo)
	http.HandleFunc("DELETE /todos/{id}", deleteTodo)

	log.Printf("Server is running on port: %s\n", port)

	if err := server.ListenAndServe(); err != nil {
		log.Fatalln(err)
	}
}

// getTodos TODOの一覧を取得する
func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	res := struct {
		Todos []Todo `json:"todos"`
	}{
		Todos: todos.List(),
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
	}
}

// addTodo TODOを追加する
func addTodo(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	req := struct {
		Name        string `json:"name"`
		Description string `json:"description"`
	}{}

	b, err := io.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		log.Println(err)
	}

	if err := json.Unmarshal(b, &req); err != nil {
		log.Println(err)
	}

	createdTodo := NewCreateTodoFactory(
		req.Name,
		req.Description,
	)

	if !todos.Add(createdTodo) {
		w.WriteHeader(http.StatusTooManyRequests)
		return
	}

	if err := json.NewEncoder(w).Encode(createdTodo); err != nil {
		log.Println(err)
	}
}

// deleteTodo TODOを削除する
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	id := r.PathValue("id")
	if todos.Get(id) == nil {
		w.WriteHeader(http.StatusNotFound)
	}
	todos = todos.Remove(id)
}

// getTodo TODOを取得する
func getTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	id := r.PathValue("id")
	todo := todos.Get(id)
	if todo == nil {
		w.WriteHeader(http.StatusNotFound)
	}

	if err := json.NewEncoder(w).Encode(todo); err != nil {
		log.Println(err)
	}
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	defer mu.Unlock()

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	req := struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		IsDone      bool   `json:"is_done"`
	}{}

	b, err := io.ReadAll(r.Body)
	r.Body.Close()
	if err != nil {
		log.Println(err)
	}

	if err := json.Unmarshal(b, &req); err != nil {
		log.Println(err)
	}

	id := r.PathValue("id")
	todo := todos.Get(id)
	if todo == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	updatedTodo := NewTodo(todo.ID, req.Name, req.Description, req.IsDone)
	if !todos.Update(updatedTodo) {
		w.WriteHeader(http.StatusNotAcceptable)
		return
	}

	if err := json.NewEncoder(w).Encode(updatedTodo); err != nil {
		log.Println(err)
	}
}
