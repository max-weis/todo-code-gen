package entity

import (
	"context"
)

type Repository interface {
	CreateTodo(context.Context, CreateTodoParams) (int, error)
	DeleteTodoById(context.Context, int64) error
	GetTodoById(context.Context, int64) (Todo, error)
	ListTodos(context.Context, ListTodosParams) ([]ListTodosRow, error)
	UpdateTodo(context.Context, UpdateTodoParams) error
}

type todoRepository struct {
	queries *Queries
}

func (t todoRepository) CreateTodo(ctx context.Context, params CreateTodoParams) (int, error) {
	todo, err := t.queries.CreateTodo(ctx, params)
	if err != nil {
		return -1, err
	}
	id, err := todo.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(id), nil
}

func (t todoRepository) DeleteTodoById(ctx context.Context, i int64) error {
	return t.queries.DeleteTodoById(ctx, i)
}

func (t todoRepository) GetTodoById(ctx context.Context, i int64) (Todo, error) {
	return t.queries.GetTodoById(ctx, i)
}

func (t todoRepository) ListTodos(ctx context.Context, params ListTodosParams) ([]ListTodosRow, error) {
	return t.queries.ListTodos(ctx, params)
}

func (t todoRepository) UpdateTodo(ctx context.Context, params UpdateTodoParams) error {
	return t.queries.UpdateTodo(ctx, params)
}

func ProvideRepository(queries *Queries) Repository {
	return todoRepository{queries: queries}
}
