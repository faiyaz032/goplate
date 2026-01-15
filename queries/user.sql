-- name: CreateUser :one
INSERT INTO users (
    username,
    email,
    password
) VALUES (
    $1, $2, $3
)
RETURNING id, username, email, password, created_at, updated_at, deleted_at;

-- name: GetUser :one
SELECT id, username, email, password, created_at, updated_at, deleted_at
FROM users
WHERE id = $1 AND deleted_at IS NULL LIMIT 1;

-- name: GetUserByEmail :one
SELECT id, username, email, password, created_at, updated_at, deleted_at
FROM users
WHERE email = $1 AND deleted_at IS NULL LIMIT 1;

-- name: ListUsers :many
SELECT id, username, email, created_at, updated_at
FROM users
WHERE deleted_at IS NULL
ORDER BY created_at DESC;

-- name: UpdateUser :one
UPDATE users
SET
    username = COALESCE(sqlc.narg('username'), username),
    email = COALESCE(sqlc.narg('email'), email),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted_at IS NULL
RETURNING id, username, email, updated_at;

-- name: SoftDeleteUser :exec
UPDATE users
SET deleted_at = CURRENT_TIMESTAMP
WHERE id = $1 AND deleted_at IS NULL;

-- name: RestoreUser :exec
UPDATE users
SET deleted_at = NULL
WHERE id = $1;
