-- name: GetTodoById :one
SELECT t.id, t.title, t.description, t.status
FROM todos AS t
WHERE id = ? LIMIT 1;

-- name: ListTodos :many
SELECT t.id, t.title
FROM todos AS t LIMIT ?
OFFSET ?;

-- name: CreateTodo :execresult
INSERT INTO todos (title, description, status)
VALUES (?, ?, ?);

-- name: UpdateTodo :exec
UPDATE todos
SET title=?,
    description=?,
    status=?
WHERE id = ? LIMIT 1;

-- name: DeleteTodoById :exec
DELETE
FROM todos
WHERE id = ?;