-- name: GetRestaurant :one
SELECT * FROM restaurant
WHERE id = $1;

-- name: CreateRestaurant :one
INSERT INTO restaurant (
 name, address, phone_number, email
) VALUES (
  $1, $2, $3, $4
)
RETURNING *;

-- name: DeleteRestaurant :exec
DELETE FROM restaurant
WHERE id = $1;
