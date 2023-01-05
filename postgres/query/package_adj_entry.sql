-- name: CreatePackageAdjEntry :one
-- description: Create a package adjustment entry
INSERT INTO package_adj_entries (package_id,
                                 amount,
                                 uom_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetPackageAdjEntry :one
-- description: Get a package adjustment entry
SELECT *
FROM package_adj_entries
WHERE id = $1
LIMIT 1;

-- name: ListPackageAdjEntries :many
-- description: List package adjustment entries by package id
SELECT *
FROM package_adj_entries
WHERE package_id = $1
ORDER BY id
LIMIT $2 OFFSET $3;
