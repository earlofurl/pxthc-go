-- name: CreateItemType :one
-- description: Create a new item type
INSERT INTO item_types (product_form,
                        product_modifier,
                        uom_default,
                        product_category_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetItemType :one
-- description: Get an item type by id
SELECT *
FROM item_types
WHERE id = $1
LIMIT 1;

-- name: ListItemTypes :many
-- description: List all item types
SELECT *
FROM item_types
ORDER BY id;

-- name: UpdateItemType :one
-- description: Update an item type by ID
UPDATE item_types
SET product_form        = COALESCE(sqlc.narg(product_form), product_form),
    product_modifier    = COALESCE(sqlc.narg(product_modifier), product_modifier),
    uom_default         = COALESCE(sqlc.narg(uom_default), uom_default),
    product_category_id = COALESCE(sqlc.narg(product_category_id), product_category_id),
    updated_at          = NOW()
WHERE id = sqlc.arg(id)
    RETURNING *;

-- name: DeleteItemType :exec
-- description: Delete an item type by ID
DELETE
FROM item_types
WHERE id = $1;
