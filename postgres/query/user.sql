-- name: CreateUser :one
INSERT INTO users (username, email, hashed_password, first_name, last_name, phone, role)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetUserByEmail :one
SELECT *
FROM users
WHERE email = $1
LIMIT 1;

-- name: GetUserByID :one
SELECT *
FROM users
WHERE id = $1
LIMIT 1;

-- name: UpdateUser :one
UPDATE users
SET hashed_password     = COALESCE(sqlc.narg(hashed_password), hashed_password),
    password_changed_at = COALESCE(sqlc.narg(password_changed_at), password_changed_at),
    first_name          = COALESCE(sqlc.narg(first_name), first_name),
    last_name           = COALESCE(sqlc.narg(last_name), last_name),
    email               = COALESCE(sqlc.narg(email), email),
    phone               = COALESCE(sqlc.narg(phone), phone),
    role                = COALESCE(sqlc.narg(role), role)
WHERE username = sqlc.arg(username)
RETURNING *;
