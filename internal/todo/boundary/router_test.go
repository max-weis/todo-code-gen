package boundary

import (
	"context"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"testing"
)

var networkName = "integrationtest"

type mysqlContainer struct {
	testcontainers.Container
	host string
	port string
}

type todoContainer struct {
	testcontainers.Container
	host string
	port string
}

func setupMySQL(ctx context.Context) (*mysqlContainer, error) {
	_, err := testcontainers.GenericNetwork(ctx, testcontainers.GenericNetworkRequest{
		NetworkRequest: testcontainers.NetworkRequest{
			Name:           networkName,
			CheckDuplicate: true,
		},
	})

	pathToSchema, err := filepath.Abs("../entity/schema.sql")
	if err != nil {
		return nil, err
	}

	pathToImport, err := filepath.Abs("../entity/import.sql")
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

	return &mysqlContainer{Container: container, host: ip, port: mappedPort.Port()}, nil
}

func setupTodoService(ctx context.Context, mysqlC *mysqlContainer) (*todoContainer, error) {
	req := testcontainers.ContainerRequest{
		Image:        "maxweis/todo-code-gen",
		ExposedPorts: []string{"8080/tcp"},
		Env: map[string]string{
			"MYSQL_USER":     "mysql",
			"MYSQL_PASSWORD": "mysql",
			"MYSQL_HOST":     mysqlC.host,
			"MYSQL_PORT":     mysqlC.port,
			"MYSQL_DB":       "mysql",
		},
		WaitingFor: wait.ForLog("{\"message\":\"⇨ http server started on [::]:8080\"}"),
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

	mappedPort, err := container.MappedPort(ctx, "8080")
	if err != nil {
		return nil, err
	}

	return &todoContainer{Container: container, host: ip, port: mappedPort.Port()}, nil
}

func TestTodoRouter_CreateTodo(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	ctx := context.Background()

	mysqlC, err := setupMySQL(ctx)
	assert.Nil(t, err, "container should be started")

	defer mysqlC.Terminate(ctx)

	todoC, err := setupTodoService(ctx, mysqlC)
	assert.Nil(t, err, "container should be started")

	defer todoC.Terminate(ctx)

	client := Client{
		Server: fmt.Sprintf("http://%s:%s", todoC.host, todoC.port),
		Client: &http.Client{},
	}

	var description string
	description = "test"
	resp, err := client.CreateTodo(context.Background(), CreateTodoJSONRequestBody{
		Description: &description,
		Done:        false,
		Title:       "test",
	})
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err, "container should be started")

	assert.Nil(t, err, "err should be nil")
	assert.Equal(t, "/todos/11", string(bodyBytes), "path should equal /api/todos/11")
}
