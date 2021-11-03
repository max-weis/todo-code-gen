// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	entity "todo-code-gen/internal/todo/entity"

	mock "github.com/stretchr/testify/mock"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// CreateTodo provides a mock function with given fields: ctx, title, description, status
func (_m *Controller) CreateTodo(ctx context.Context, title string, description string, status bool) (int, error) {
	ret := _m.Called(ctx, title, description, status)

	var r0 int
	if rf, ok := ret.Get(0).(func(context.Context, string, string, bool) int); ok {
		r0 = rf(ctx, title, description, status)
	} else {
		r0 = ret.Get(0).(int)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string, bool) error); ok {
		r1 = rf(ctx, title, description, status)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteTodoById provides a mock function with given fields: ctx, id
func (_m *Controller) DeleteTodoById(ctx context.Context, id int) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetTodoById provides a mock function with given fields: ctx, id
func (_m *Controller) GetTodoById(ctx context.Context, id int) (*entity.Todo, error) {
	ret := _m.Called(ctx, id)

	var r0 *entity.Todo
	if rf, ok := ret.Get(0).(func(context.Context, int) *entity.Todo); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Todo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTodos provides a mock function with given fields: ctx, limit, offset
func (_m *Controller) ListTodos(ctx context.Context, limit int, offset int) (*[]entity.ListTodosRow, error) {
	ret := _m.Called(ctx, limit, offset)

	var r0 *[]entity.ListTodosRow
	if rf, ok := ret.Get(0).(func(context.Context, int, int) *[]entity.ListTodosRow); ok {
		r0 = rf(ctx, limit, offset)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*[]entity.ListTodosRow)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int, int) error); ok {
		r1 = rf(ctx, limit, offset)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateTodo provides a mock function with given fields: ctx, id, title, description, status
func (_m *Controller) UpdateTodo(ctx context.Context, id int, title string, description string, status bool) error {
	ret := _m.Called(ctx, id, title, description, status)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int, string, string, bool) error); ok {
		r0 = rf(ctx, id, title, description, status)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}