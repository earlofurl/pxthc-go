-- name: CreateProductCategory :one
-- description: Create a new product category
INSERT INTO product_categories (name)
VALUES ($1)
RETURNING *;

-- name: GetProductCategory :one
-- description: Get a product category by ID
SELECT *
FROM product_categories
WHERE id = $1
LIMIT 1;

-- name: GetProductCategoryByName :one
-- description: Get a product category by name
SELECT *
FROM product_categories
WHERE name = $1
LIMIT 1;

-- name: ListProductCategories :many
-- description: List all product categories
SELECT *
FROM product_categories
ORDER BY name;

-- name: UpdateProductCategory :one
-- description: Update a product category
UPDATE product_categories
SET name       = $2,
    updated_at = NOW()
WHERE id = $1
RETURNING *;

-- name: DeleteProductCategory :exec
-- description: Delete a product category
DELETE
FROM product_categories
WHERE id = $1;