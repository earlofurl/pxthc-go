-- name: CreatePackageAdjustment :one
-- description: Create a package adjustment
INSERT INTO package_adjustments (from_package_id,
                                 to_package_id,
                                 amount,
                                 uom_id)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetPackageAdjustment :one
-- description: Get a package adjustment by id
SELECT *
FROM package_adjustments
WHERE id = $1
LIMIT 1;

-- name: ListPackageAdjustments :many
-- description: List package adjustments
SELECT *
FROM package_adjustments
WHERE from_package_id = $1
   OR to_package_id = $2
ORDER BY id
LIMIT $3 OFFSET $4;
