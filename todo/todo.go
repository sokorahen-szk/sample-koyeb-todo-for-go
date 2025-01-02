package todo

import "github.com/google/uuid"

type Todo struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
}

func NewTodo(id, name, description string, isDone bool) Todo {
	return Todo{
		ID:          id,
		Name:        name,
		Description: description,
		IsDone:      isDone,
	}
}

func NewCreateTodoFactory(name, description string) Todo {
	return NewTodo(uuid.NewString(), name, description, false)
}

type TodoList struct {
	items         []Todo
	maxTodoLength int
}

type TodoListOption func(*TodoList)

func NewTodoList(options ...TodoListOption) *TodoList {
	defaultTodoMaxLimit := 10
	tl := &TodoList{
		items:         make([]Todo, 0, defaultTodoMaxLimit),
		maxTodoLength: defaultTodoMaxLimit,
	}

	for _, opt := range options {
		opt(tl)
	}

	return tl
}

func (tl *TodoList) Count() int {
	return len(tl.items)
}

func (tl *TodoList) Add(todo Todo) bool {
	if tl.Count() <= tl.maxTodoLength {
		tl.items = append(tl.items, todo)

		return true
	}

	return false
}

func (tl *TodoList) Remove(id string) *TodoList {
	for idx, todo := range tl.items {
		if todo.ID == id {
			return NewTodoList(TodoListOption(func(t *TodoList) {
				t.items = tl.items[:idx+copy(tl.items[idx:], tl.items[idx+1:])]
			}))
		}
	}

	return NewTodoList(TodoListOption(func(t *TodoList) {
		t.items = tl.items
	}))
}

func (tl *TodoList) Update(todo Todo) bool {
	for idx, t := range tl.items {
		if t.ID == todo.ID {
			tl.items[idx] = todo

			return true
		}
	}

	return false
}

func (tl *TodoList) Get(id string) *Todo {
	for _, todo := range tl.items {
		if todo.ID == id {
			return &todo
		}
	}

	return nil
}

func (tl *TodoList) List() []Todo {
	unfinishedTodos := make([]Todo, 0)
	for _, todo := range tl.items {
		// 完了済みはTODOリストから除外される
		if todo.IsDone {
			continue
		}

		unfinishedTodos = append(unfinishedTodos, todo)
	}

	return unfinishedTodos
}
