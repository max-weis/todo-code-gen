package boundary

import (
	"database/sql"
	"github.com/stretchr/testify/assert"
	"testing"
	"todo-code-gen/internal/todo/entity"
)

func Test_todosToList(t *testing.T) {
	rows := []entity.ListTodosRow{
		{ID: 1, Title: "test1"},
		{ID: 2, Title: "test2"},
		{ID: 3, Title: "test3"},
	}

	todos := todosToList(rows)

	assert.Equal(t, 3, len(*todos), "length should equal 3")
}

func Test_toListDTO(t *testing.T) {
	todo := entity.Todo{
		ID:    1,
		Title: "test",
		Description: sql.NullString{
			String: "test",
			Valid:  true,
		},
		Status: 1,
	}

	listDTO := toListDTO(todo)

	assert.Equal(t, 1, listDTO.Id, "id should equal 1")
	assert.Equal(t, "test", listDTO.Title, "title should equal test")
	assert.True(t, listDTO.Done, "done should be true")
}

func Test_toFullDTO(t *testing.T) {
	todo := entity.Todo{
		ID:    1,
		Title: "test",
		Description: sql.NullString{
			String: "test",
			Valid:  true,
		},
		Status: 1,
	}

	fullDTO := toFullDTO(todo)

	assert.Equal(t, 1, fullDTO.Id, "id should equal 1")
	assert.Equal(t, "test", fullDTO.Title, "title should equal test")
	assert.Equal(t, "test", *fullDTO.Description, "description should equal test")
	assert.True(t, fullDTO.Done, "done should be true")
}

func Test_toBaseDTO(t *testing.T) {
	todo := entity.Todo{
		ID:    1,
		Title: "test",
		Description: sql.NullString{
			String: "test",
			Valid:  true,
		},
		Status: 1,
	}

	baseDTO := toBaseDTO(todo)

	assert.Equal(t, "test", baseDTO.Title, "title should equal test")
	assert.Equal(t, "test", *baseDTO.Description, "description should equal test")
	assert.True(t, baseDTO.Done, "done should be true")
}

func Test_stateToBool(t *testing.T) {
	assert.False(t, stateToBool(0), "0 should be false")
	assert.True(t, stateToBool(1), "1 should be true")
}
