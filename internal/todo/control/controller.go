package control

import (
	"context"
	"database/sql"
	"github.com/labstack/gommon/log"
	"todo-code-gen/internal/todo/entity"
)

type Controller interface {
	GetTodoById(ctx context.Context, id int) (*entity.Todo, error)
	ListTodos(ctx context.Context, limit, offset int) (*[]entity.ListTodosRow, error)
	CreateTodo(ctx context.Context, title, description string, status bool) (int, error)
	UpdateTodo(ctx context.Context, id int, title, description string, status bool) error
	DeleteTodoById(ctx context.Context, id int) error
}

type TodoControllerImpl struct {
	repository entity.Repository
}

func ProvideController(repository entity.Repository) *TodoControllerImpl {
	return &TodoControllerImpl{repository: repository}
}

func (t *TodoControllerImpl) GetTodoById(ctx context.Context, id int) (*entity.Todo, error) {
	log.Infof("Get todo with id: %d", id)
	todo, err := t.repository.GetTodoById(ctx, int64(id))
	if err != nil {
		log.Warnf("Could not find todo with id: %d", id)
		return nil, err
	}

	return &todo, nil
}

func (t *TodoControllerImpl) ListTodos(ctx context.Context, limit, offset int) (*[]entity.ListTodosRow, error) {
	log.Infof("Get todos with limit: %d, offset: %s", limit, offset)
	todos, err := t.repository.ListTodos(ctx, entity.ListTodosParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		log.Warnf("Could not find todos with limit: %d, offset: %s", limit, offset)
		return nil, err
	}

	return &todos, nil
}

func (t *TodoControllerImpl) CreateTodo(ctx context.Context, title, description string, status bool) (int, error) {
	log.Infof("Create todo with title: %s, description: %s, status: %t", title, description, status)
	id, err := t.repository.CreateTodo(ctx, entity.CreateTodoParams{
		Title: title,
		Description: sql.NullString{
			String: description,
			Valid:  true,
		},
		Status: stateToInt(status),
	})
	if err != nil {
		log.Warnf("Could not create todo")
		return -1, err
	}

	return id, nil
}

func (t *TodoControllerImpl) UpdateTodo(ctx context.Context, id int, title, description string, status bool) error {
	log.Infof("Update todo with id: %d", id)
	err := t.repository.UpdateTodo(ctx, entity.UpdateTodoParams{
		Title: title,
		Description: sql.NullString{
			String: description,
			Valid:  true,
		},
		Status: stateToInt(status),
		ID:     int64(id),
	})
	if err != nil {
		log.Warnf("Could not update todo with id: %d", id)
		return err
	}

	return nil
}

func (t *TodoControllerImpl) DeleteTodoById(ctx context.Context, id int) error {
	log.Infof("Delete todo with id: %d", id)
	return t.repository.DeleteTodoById(ctx, int64(id))
}

func stateToInt(state bool) int32 {
	if state {
		return 1
	}

	return 0
}
