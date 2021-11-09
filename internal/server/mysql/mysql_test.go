package mysql

import (
	"context"
	"database/sql"
	"github.com/docker/go-connections/nat"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"os"
	"path/filepath"
	"testing"
	"todo-code-gen/internal/config"
	"todo-code-gen/internal/todo/entity"
)

var (
	port   nat.Port
)

func TestMain(m *testing.M) {
	ctx := context.Background()

	container := setupMySQL(ctx)
	code := m.Run()
	container.Terminate(ctx)

	os.Exit(code)
}

func setupMySQL(ctx context.Context) testcontainers.Container {
	pathToSchema, _ := filepath.Abs("../../todo/entity/schema.sql")

	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
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
		},
		Started: true,
	})
	if err != nil {
		log.Printf("Could not create container, %+v", err)
		os.Exit(1)
	}

	port, err = container.MappedPort(ctx, "3306")
	if err != nil {
		log.Printf("Could not get port 3306, %+v", err)
		os.Exit(1)
	}

	return container
}

func Test_ProvideDatabase(t *testing.T) {
	databaseConfig := config.DatabaseConfig{
		User:     "mysql",
		Password: "mysql",
		Host:     "localhost",
		Port:     port.Port(),
		DbName:   "mysql",
	}

	db := ProvideDatabase(databaseConfig)

	_, err := db.CreateTodo(context.Background(), entity.CreateTodoParams{
		Title:       "test",
		Description: sql.NullString{},
		Status:      0,
	})

	assert.NotNil(t, db, "db should not be nil")
	assert.Nil(t, err, "error should be nil")
}
