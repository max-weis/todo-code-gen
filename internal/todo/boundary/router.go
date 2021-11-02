package boundary

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"todo-code-gen/internal/todo/control"
)

type TodoRouter struct {
	controller control.TodoController
}

func ProvideRouter(controller control.TodoController) *TodoRouter {
	return &TodoRouter{controller: controller}
}

func (t *TodoRouter) RegisterHandlersWithBaseURL(router EchoRouter, baseURL string) {
	RegisterHandlersWithBaseURL(router, t, baseURL)
}

func (t *TodoRouter) GetTodos(ctx echo.Context, params GetTodosParams) error {
	todos, err := t.controller.ListTodos(*params.Limit, *params.Offset)
	if err != nil {
		log.Errorf("could not get dto: %+v", err)
		return err
	}

	dto := todosToList(*todos)

	if err = ctx.JSON(http.StatusOK, dto); err != nil {
		log.Errorf("could not marshal todo: %+v", err)
		return err
	}

	return nil
}

func (t *TodoRouter) CreateTodo(ctx echo.Context) error {
	var base *TodoBase = &TodoBase{}
	if err := ctx.Bind(base); err != nil {
		log.Errorf("could not parse request body: %+v", err)
		return err
	}

	id, err := t.controller.CreateTodo(base.Title, *base.Description, base.Done)
	if err != nil {
		log.Errorf("could not create todo: %+v", err)
		return err
	}

	payload := fmt.Sprintf("/todo/%d", id)

	if err = ctx.String(201, payload); err != nil {
		log.Errorf("could not marshal todo: %+v", err)
		return err
	}

	return nil
}

func (t *TodoRouter) DeleteTodo(ctx echo.Context, todoId int) error {
	err := t.controller.DeleteTodoById(todoId)
	if err != nil {
		log.Errorf("could not remove todo: %+v", err)
		return err
	}

	return nil
}

func (t *TodoRouter) GetTodo(ctx echo.Context, todoId int) error {
	todo, err := t.controller.GetTodoById(todoId)
	if err != nil {
		log.Errorf("could not get todo: %+v", err)
		return err
	}

	dto := toFullDTO(*todo)

	if err = ctx.JSON(http.StatusOK, dto); err != nil {
		log.Errorf("could not marshal todo: %+v", err)
		return err
	}

	return nil
}

func (t *TodoRouter) UpdateTodo(ctx echo.Context, todoId int) error {
	var base TodoBase
	if err := ctx.Bind(base); err != nil {
		log.Errorf("could not parse request body: %+v", err)
		return err
	}

	err := t.controller.UpdateTodo(todoId, base.Title, *base.Description, base.Done)
	if err != nil {
		log.Errorf("could not update todo: %+v", err)
		return err
	}

	return nil
}
