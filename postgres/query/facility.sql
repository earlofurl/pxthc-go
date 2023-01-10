-- name: CreateFacility :one
-- description: Create a new facility
INSERT INTO facilities (name, license_number)
VALUES ($1, $2)
RETURNING *;

-- name: GetFacilityByID :one
-- description: Get a facility by ID
SELECT *
FROM facilities
WHERE id = $1
LIMIT 1;

-- name: GetFacilityByName :one
-- description: Get a facility by name
SELECT *
FROM facilities
WHERE name ILIKE $1
LIMIT 1;

-- name: ListFacilities :many
-- description: List all facilities
SELECT *
FROM facilities;

-- name: UpdateFacility :one
-- description: Update a facility by ID
UPDATE facilities
SET name           = COALESCE(sqlc.narg(name), name),
    license_number = COALESCE(sqlc.narg(license_number), license_number)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteFacility :exec
-- description: Delete a facility by ID
DELETE
FROM facilities
WHERE id = $1;
