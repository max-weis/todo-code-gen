//go:build integration
// +build integration

package mysql

import (
	"context"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"path/filepath"
	"testing"
	"todo-code-gen/internal/config"
	"todo-code-gen/internal/todo/entity"
)

type mysqlContainer struct {
	testcontainers.Container
	host string
	port string
}

func setupMySQL(ctx context.Context) (*mysqlContainer, error) {
	pathToSchema, err := filepath.Abs("../../todo/entity/schema.sql")
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
			pathToSchema: "/docker-entrypoint-initdb.d/init.sql",
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

	return &mysqlContainer{Container: container, host: ip, port: mappedPort.Port()}, nil
}

func Test_ProvideDatabase(t *testing.T) {
	ctx := context.Background()

	mysqlC, err := setupMySQL(ctx)
	if err != nil {
		t.Fatal(err)
	}

	defer mysqlC.Terminate(ctx)

	databaseConfig := config.DatabaseConfig{
		User:     "mysql",
		Password: "mysql",
		Host:     mysqlC.host,
		Port:     mysqlC.port,
		DbName:   "mysql",
	}

	db := ProvideDatabase(databaseConfig)

	_, err = db.CreateTodo(context.Background(), entity.CreateTodoParams{
		Title:       "test",
		Description: sql.NullString{},
		Status:      0,
	})

	assert.Nil(t, err, "error should be nil")
}
