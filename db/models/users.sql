-- name: CreateUser :one
INSERT INTO users (
    created_at, updated_at, first_name, last_name, email, picture_url
) VALUES (
    now(), now(), $1, $2, $3, $4
)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1 LIMIT 1;
