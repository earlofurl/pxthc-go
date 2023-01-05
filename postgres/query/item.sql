-- name: CreateItem :one
-- description: Create a new item
INSERT INTO items (description, is_used, item_type_id, strain_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetItem :one
-- description: Get an item by ID
SELECT *
FROM items
WHERE id = $1
LIMIT 1;

-- name: ListItems :many
-- description: List all items
SELECT i.*, s.name AS strain_name, t.product_form AS product_form, t.product_modifier AS product_modifier
FROM items i
         INNER JOIN strains s ON i.strain_id = s.id
         INNER JOIN item_types t ON i.item_type_id = t.id
ORDER BY strain_id;

-- name: UpdateItem :one
-- description: Update an item by ID
UPDATE items
SET description  = COALESCE(sqlc.narg(description), description),
    is_used      = COALESCE(sqlc.narg(is_used), is_used),
    item_type_id = COALESCE(sqlc.narg(item_type_id), item_type_id),
    strain_id    = COALESCE(sqlc.narg(strain_id), strain_id)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteItem :exec
-- description: Delete an item by ID
DELETE
FROM items
WHERE id = $1;
