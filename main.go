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
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

const MAX_TODO_LENGTH = 10

// NOTE: 簡略化のためメモリ上に保存する
var todos = make([]Todo, 0, MAX_TODO_LENGTH)

func main() {
	port := os.Getenv("APP_PORT")
	server := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: nil,
	}
	http.HandleFunc("GET /todos", getTodos)
	http.HandleFunc("POST /todos", addTodo)
	// TODO: 削除機能を追加する http.HandleFunc("DELETE /todos/{id}", deleteTodo)
	// TODO: 更新機能を追加する http.HandleFunc("PUT /todos/{id}", updateTodo)
	// TODO: IDを指定して取得する機能を追加する http.HandleFunc("GET /todos/{id}", getTodo)

	log.Printf("Server is running on port: %s\n", port)
	server.ListenAndServe()
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

	todos = append(todos, Todo{
		ID:          len(todos) + 1,
		Name:        req.Name,
		Description: req.Description,
	})

	return
}
