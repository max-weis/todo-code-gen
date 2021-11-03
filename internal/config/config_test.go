//go:build unit
// +build unit

package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestReadEnvVarWithDefaults(t *testing.T) {
	assert.Equal(t, "mysql", readEnvVar("MYSQL_USER", "mysql"), "they should be equal")
	assert.Equal(t, "mysql", readEnvVar("MYSQL_PASSWORD", "mysql"), "they should be equal")
	assert.Equal(t, "mysql", readEnvVar("MYSQL_HOST", "mysql"), "they should be equal")
	assert.Equal(t, "mysql", readEnvVar("MYSQL_PORT", "mysql"), "they should be equal")
	assert.Equal(t, "mysql", readEnvVar("MYSQL_DB", "mysql"), "they should be equal")
}

func TestReadEnvVarWithValues(t *testing.T) {
	os.Setenv("MYSQL_USER", "test")
	defer os.Unsetenv("MYSQL_USER")
	os.Setenv("MYSQL_PASSWORD", "test")
	defer os.Unsetenv("MYSQL_PASSWORD")
	os.Setenv("MYSQL_HOST", "test")
	defer os.Unsetenv("MYSQL_HOST")
	os.Setenv("MYSQL_PORT", "test")
	defer os.Unsetenv("MYSQL_PORT")
	os.Setenv("MYSQL_DB", "test")
	defer os.Unsetenv("MYSQL_DB")

	assert.Equal(t, "test", readEnvVar("MYSQL_USER", ""), "they should be equal")
	assert.Equal(t, "test", readEnvVar("MYSQL_PASSWORD", ""), "they should be equal")
	assert.Equal(t, "test", readEnvVar("MYSQL_HOST", ""), "they should be equal")
	assert.Equal(t, "test", readEnvVar("MYSQL_PORT", ""), "they should be equal")
	assert.Equal(t, "test", readEnvVar("MYSQL_DB", ""), "they should be equal")
}

func TestNewDatabaseConfig(t *testing.T) {
	config := NewDatabaseConfig()

	assert.Equal(t, "mysql", config.User, "they should be equal")
	assert.Equal(t, "mysql", config.Password, "they should be equal")
	assert.Equal(t, "localhost", config.Host, "they should be equal")
	assert.Equal(t, "3306", config.Port, "they should be equal")
	assert.Equal(t, "todo", config.DbName, "they should be equal")
}
