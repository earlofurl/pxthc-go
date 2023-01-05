-- name: CreatePackageTag :one
-- description: Create a package tag
INSERT INTO package_tags (tag_number, is_assigned, is_provisional, is_active, assigned_package_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetPackageTag :one
-- description: Get a package tag by ID
SELECT *
FROM package_tags
WHERE id = $1
LIMIT 1;

-- name: GetPackageTagByTagNumber :one
-- description: Get a package tag by tag number
SELECT *
FROM package_tags
WHERE tag_number = $1
LIMIT 1;

-- name: ListPackageTags :many
-- description: List package tags by limit and offset
SELECT *
FROM package_tags
WHERE is_assigned = $1
ORDER BY id
LIMIT $2 OFFSET $3;

-- name: UpdatePackageTag :one
-- description: Update a package tag
UPDATE package_tags
SET is_assigned         = COALESCE(sqlc.narg(is_assigned), is_assigned),
    is_provisional      = COALESCE(sqlc.narg(is_provisional), is_provisional),
    is_active           = COALESCE(sqlc.narg(is_active), is_active),
    assigned_package_id = COALESCE(sqlc.narg(assigned_package_id), assigned_package_id)
WHERE id = sqlc.arg(id)
RETURNING *;
