-- name: CreateTodo :one

INSERT INTO
    "todos" (id, title, description, status)
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: GetTodo :one

SELECT * FROM "todos" WHERE id = $1;

-- name: GetTodos :many

SELECT * FROM "todos";

-- name: UpdateTodo :one

UPDATE "todos"
SET
    title = COALESCE(sqlc.narg('title'), title),
    description = COALESCE(
        sqlc.narg('description'),
        description
    ),
    status = COALESCE(sqlc.narg('status'), status),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 RETURNING *;

-- name: UpdateTodoStatus :one

UPDATE "todos" SET status = $2 WHERE id = $1 RETURNING *;

-- name: DeleteTodo :one

DELETE FROM "todos" WHERE id = $1 RETURNING *;