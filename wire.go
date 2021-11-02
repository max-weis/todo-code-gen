//go:build wireinject
// +build wireinject

//go:generate wire .
package main

import (
	"github.com/google/wire"
	"todo-code-gen/internal/config"
	"todo-code-gen/internal/server/mysql"
	"todo-code-gen/internal/todo/boundary"
	"todo-code-gen/internal/todo/control"
	"todo-code-gen/internal/todo/entity"
)

type AppContext struct {
	*boundary.TodoRouter
}

func Initialize() *AppContext {
	panic(wire.Build(
		// Database
		config.NewDatabaseConfig,
		mysql.ProvideDatabase,

		// Todo
		entity.ProvideRepository,
		control.ProvideController,
		wire.Bind(new(control.TodoController), new(*control.TodoControllerImpl)),

		boundary.ProvideRouter,

		wire.Struct(new(AppContext), "*"),
	))
}
