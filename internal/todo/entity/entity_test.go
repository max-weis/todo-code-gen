package entity

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"
	"path/filepath"
	"testing"
)

var mysqlC mysqlContainer

type mysqlContainer struct {
	testcontainers.Container
	host    string
	port    string
	queries Queries
}

func TestMain(m *testing.M) {
	ctx := context.Background()
	mysqlC, err := setupMySQL(ctx)
	if err != nil {
		os.Exit(1)
	}
	code := m.Run()
	mysqlC.Terminate(ctx)
	os.Exit(code)
}

func setupMySQL(ctx context.Context) (*mysqlContainer, error) {
	pathToSchema, err := filepath.Abs("./schema.sql")
	if err != nil {
		return nil, err
	}

	pathToImport, err := filepath.Abs("./import.sql")
	if err != nil {
		return nil, err
	}

	req := testcontainers.ContainerRequest{
		Image:        "mysql",
		ExposedPorts: []string{"3306/tcp"},
		Env: map[string]string{
			"MYSQL_USER":          "mysql",
			"MYSQL_PASSWORD":      "mysql",
			"MYSQL_ROOT_PASSWORD": "mysql",
			"MYSQL_DATABASE":      "mysql",
		},
		BindMounts: map[string]string{
			pathToSchema: "/docker-entrypoint-initdb.d/1-schema.sql",
			pathToImport: "/docker-entrypoint-initdb.d/2-import.sql",
		},
		WaitingFor: wait.ForLog("port: 3306  MySQL Community Server - GPL"),
	}

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if err != nil {
		return nil, err
	}

	ip, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}

	mappedPort, err := container.MappedPort(ctx, "3306")
	if err != nil {
		return nil, err
	}

	dsn := fmt.Sprintf("mysql:mysql@tcp(%s:%s)/mysql?charset=utf8mb4&parseTime=True&loc=Local&tls=preferred&timeout=5s", ip, mappedPort.Port())
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Err(err).Msg("Could not connect to mysql")
	}

	queries := New(db)

	return &mysqlContainer{Container: container, host: ip, port: mappedPort.Port(), queries: *queries}, nil
}

func TestTodoRepository_CreateTodo(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	repository := todoRepository{queries: &mysqlC.queries}
	id, err := repository.CreateTodo(context.Background(), CreateTodoParams{
		Title: "test",
		Description: sql.NullString{
			String: "test",
			Valid:  true,
		},
		Status: 0,
	})

	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, 11, id, "id should equal 1")
}

func TestTodoRepository_DeleteTodoById(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	repository := todoRepository{queries: &mysqlC.queries}
	err := repository.DeleteTodoById(context.Background(), 1)

	assert.Nil(t, err, "err should be nil")
}

func TestTodoRepository_GetTodoById(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	repository := todoRepository{queries: &mysqlC.queries}
	todo, err := repository.GetTodoById(context.Background(), 1)

	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, int64(1), todo.ID, "id should be equal")
	assert.Equal(t, "ligula", todo.Title, "title should be equal")
	assert.Equal(t, "orci pede venenatis non sodales", todo.Description.String, "description should be equal")
	assert.Equal(t, int32(0), todo.Status, "status should be equal")
}

func TestTodoRepository_UpdateTodo(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	repository := todoRepository{queries: &mysqlC.queries}
	err := repository.UpdateTodo(context.Background(), UpdateTodoParams{
		Title: "test",
		Description: sql.NullString{
			String: "test",
			Valid:  true,
		},
		Status: 0,
		ID:     1,
	})

	todo, err := repository.GetTodoById(context.Background(), 1)
	assert.Nil(t, err, "err should be nil")

	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, int64(1), todo.ID, "id should be equal")
	assert.Equal(t, "test", todo.Title, "title should be equal")
	assert.Equal(t, "test", todo.Description.String, "description should be equal")
	assert.Equal(t, int32(0), todo.Status, "status should be equal")
}

func TestTodoRepository_ListTodos(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	repository := todoRepository{queries: &mysqlC.queries}
	todos, err := repository.ListTodos(context.Background(), ListTodosParams{
		Limit:  5,
		Offset: 0,
	})

	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, 5, len(todos), "lenght should equal 5")
}
