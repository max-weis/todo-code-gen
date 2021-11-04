package control

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
	"todo-code-gen/internal/todo/entity"
	"todo-code-gen/mocks"
)

func TestTodoControllerImpl_GetTodoById(t *testing.T) {
	mockRepository := new(mocks.Repository)
	mockRepository.On("GetTodoById", nil, int64(1)).Return(entity.Todo{
		ID:    1,
		Title: "test",
		Description: sql.NullString{
			String: "test",
			Valid:  true,
		},
		Status: int32(0),
	}, nil)

	controller := ProvideController(mockRepository)
	todo, err := controller.GetTodoById(nil, 1)

	mockRepository.AssertExpectations(t)

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, int64(1), todo.ID, "id should be 1")
	assert.Equal(t, "test", todo.Title, "title should be test")
	assert.Equal(t, "test", todo.Description.String, "description should be test")
	assert.Equal(t, int32(0), todo.Status, "status should be 0")
}

func TestTodoControllerImpl_ListTodos(t *testing.T) {
	list := []entity.ListTodosRow{
		entity.ListTodosRow{ID: 1, Title: "test1"},
		entity.ListTodosRow{ID: 2, Title: "test2"},
		entity.ListTodosRow{ID: 3, Title: "test3"},
	}

	mockRepository := new(mocks.Repository)
	mockRepository.On("ListTodos", nil, entity.ListTodosParams{
		Limit:  10,
		Offset: 0,
	}).Return(list, nil)

	controller := ProvideController(mockRepository)
	todos, err := controller.ListTodos(nil, 10, 0)

	mockRepository.AssertExpectations(t)

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 3, len(*todos), "length should be 3")
}

func TestTodoControllerImpl_DeleteTodoById(t *testing.T) {
	mockRepository := new(mocks.Repository)
	mockRepository.On("DeleteTodoById", nil, int64(1)).Return(nil)

	controller := ProvideController(mockRepository)
	err := controller.DeleteTodoById(nil, 1)

	mockRepository.AssertExpectations(t)

	assert.Nil(t, err, "Error should be nil")
}

func TestTodoControllerImpl_UpdateTodo(t *testing.T) {
	mockRepository := new(mocks.Repository)
	mockRepository.On("UpdateTodo", nil, entity.UpdateTodoParams{
		Title: "test",
		Description: sql.NullString{
			String: "test",
			Valid:  true,
		},
		Status: int32(0),
		ID:     1,
	}).Return(nil)

	controller := ProvideController(mockRepository)
	err := controller.UpdateTodo(nil, 1, "test", "test", false)

	mockRepository.AssertExpectations(t)

	assert.Nil(t, err, "Error should be nil")
}

func TestTodoControllerImpl_CreateTodo(t *testing.T) {
	mockRepository := new(mocks.Repository)
	mockRepository.On("CreateTodo", nil, entity.CreateTodoParams{
		Title: "test",
		Description: sql.NullString{
			String: "test",
			Valid:  true,
		},
		Status: int32(0),
	}).Return(1, nil)

	controller := ProvideController(mockRepository)
	id, err := controller.CreateTodo(nil, "test", "test", false)

	mockRepository.AssertExpectations(t)

	assert.Nil(t, err, "Error should be nil")
	assert.Equal(t, 1, id, "id should be 1")
}
