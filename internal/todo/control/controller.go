package control

import (
	"context"
	"database/sql"
	"github.com/labstack/gommon/log"
	"todo-code-gen/internal/todo/entity"
)

type TodoController interface {
	GetTodoById(id int) (*entity.Todo, error)
	ListTodos(limit, offset int) (*[]entity.ListTodosRow, error)
	CreateTodo(title, description string, status bool) (int, error)
	UpdateTodo(id int, title, description string, status bool) error
	DeleteTodoById(id int) error
}

type TodoControllerImpl struct {
	repository entity.Repository
}

func ProvideController(repository entity.Repository) *TodoControllerImpl {
	return &TodoControllerImpl{repository: repository}
}

func (t *TodoControllerImpl) GetTodoById(id int) (*entity.Todo, error) {
	log.Infof("Get todo with id: %d", id)
	todo, err := t.repository.GetTodoById(context.Background(), int64(id))
	if err != nil {
		log.Warnf("Could not find todo with id: %d", id)
		return nil, err
	}

	return &todo, nil
}

func (t *TodoControllerImpl) ListTodos(limit, offset int) (*[]entity.ListTodosRow, error) {
	log.Infof("Get todos with limit: %d, offset: %s", limit, offset)
	todos, err := t.repository.ListTodos(context.Background(), entity.ListTodosParams{
		Limit:  int32(limit),
		Offset: int32(offset),
	})
	if err != nil {
		log.Warnf("Could not find todos with limit: %d, offset: %s", limit, offset)
		return nil, err
	}

	return &todos, nil
}

func (t *TodoControllerImpl) CreateTodo(title, description string, status bool) (int, error) {
	log.Infof("Create todo with title: %s, description: %s, status: %t", title, description, status)
	result, err := t.repository.CreateTodo(context.Background(), entity.CreateTodoParams{
		Title: title,
		Description: sql.NullString{
			String: description,
		},
		Status: stateToInt(status),
	})
	if err != nil {
		log.Warnf("Could not create todo")
		return -1, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Warnf("Could not get last inserted id")
		return -1, err
	}

	return int(id), nil
}

func (t *TodoControllerImpl) UpdateTodo(id int, title, description string, status bool) error {
	log.Infof("Update todo with id: %d", id)
	err := t.repository.UpdateTodo(context.Background(), entity.UpdateTodoParams{
		Title: title,
		Description: sql.NullString{
			String: description,
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

func (t *TodoControllerImpl) DeleteTodoById(id int) error {
	log.Infof("Delete todo with id: %d", id)
	return t.repository.DeleteTodoById(context.Background(), int64(id))
}

func stateToInt(state bool) int32 {
	if state {
		return 1
	}

	return 0
}
