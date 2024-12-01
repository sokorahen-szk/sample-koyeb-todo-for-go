package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Todo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

const MAX_TODO_LENGTH = 10

// NOTE: 簡略化のためメモリ上に保存する
var todos = make([]Todo, 0, MAX_TODO_LENGTH)

var incrementCounter = 0

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
		Todos: todos,
	}

	if err := json.NewEncoder(w).Encode(res); err != nil {
		log.Println(err)
	}

	return
}

// addTodo TODOを追加する
func addTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if len(todos) >= MAX_TODO_LENGTH {
		w.WriteHeader(http.StatusTooManyRequests)
		return
	}

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

	incrementCounter = incrementCounter + 1
	todos = append(todos, Todo{
		ID:          fmt.Sprintf("%d", incrementCounter),
		Name:        req.Name,
		Description: req.Description,
	})

	return
}

// deleteTodo TODOを削除する
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	id := r.PathValue("id")
	for idx, todo := range todos {
		if todo.ID == id {
			todos = todos[:idx+copy(todos[idx:], todos[idx+1:])]

			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	return
}

// getTodo TODOを取得する
func getTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	id := r.PathValue("id")
	for _, todo := range todos {
		if todo.ID == id {
			if err := json.NewEncoder(w).Encode(todo); err != nil {
				log.Println(err)
			}

			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	return
}

func updateTodo(w http.ResponseWriter, r *http.Request) {
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

	id := r.PathValue("id")
	for idx, todo := range todos {
		if todo.ID == id {
			todos[idx] = Todo{
				ID:          todo.ID,
				Name:        req.Name,
				Description: req.Description,
			}
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	return
}
