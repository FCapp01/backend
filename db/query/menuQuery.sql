-- name: GetMenu :many
SELECT * FROM menu
WHERE restaurant_id = $1;

-- name: CreateMenu :one
INSERT INTO menu (
  restaurant_id, name, description
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: DeleteMenu :exec
DELETE FROM menu
WHERE id = $1;
