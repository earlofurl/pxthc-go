-- name: CreateFacilityLocation :one
-- description: Create a new location within a facility
INSERT INTO facility_locations (name, facility_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetFacilityLocationByID :one
-- description: Get a location within a facility by ID
SELECT *
FROM facility_locations
WHERE id = $1
LIMIT 1;

-- name: GetFacilityLocationByName :one
-- description: Get a location within a facility by name
SELECT *
FROM facility_locations
WHERE name ILIKE $1
LIMIT 1;

-- name: ListFacilityLocations :many
-- description: List all locations within facilities
SELECT *
FROM facility_locations;

-- name: ListFacilityLocationsByFacility :many
-- description: List all locations within a facility
SELECT *
FROM facility_locations
WHERE facility_id = $1;

-- name: UpdateFacilityLocation :one
-- description: Update a location within a facility
UPDATE facility_locations
SET name        = COALESCE(sqlc.narg(name), name),
    facility_id = COALESCE(sqlc.narg(facility_id), facility_id)
WHERE id = sqlc.arg(id)
RETURNING *;

-- name: DeleteFacilityLocation :exec
-- description: Delete a location within a facility
DELETE
FROM facility_locations
WHERE id = $1;
