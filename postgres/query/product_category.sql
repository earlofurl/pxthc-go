-- name: GetProductCategoryByID :one
-- description: Get a product category by ID
SELECT *
FROM product_categories
WHERE id = $1
LIMIT 1;

-- name: GetProductCategoryByName :one
-- description: Get a product category by name
SELECT *
FROM product_categories
WHERE name ILIKE $1
LIMIT 1;

-- name: ListProductCategories :many
-- description: List all product categories
SELECT *
FROM product_categories
ORDER BY name;
