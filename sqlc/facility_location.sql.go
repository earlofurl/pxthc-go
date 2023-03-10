// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: facility_location.sql

package sqlc

import (
	"context"

	"github.com/gobuffalo/nulls"
)

const createFacilityLocation = `-- name: CreateFacilityLocation :one
INSERT INTO facility_locations (name, facility_id)
VALUES ($1, $2)
RETURNING id, created_at, updated_at, name, facility_id
`

type CreateFacilityLocationParams struct {
	Name       string `json:"name"`
	FacilityID int64  `json:"facility_id"`
}

// description: Create a new location within a facility
func (q *Queries) CreateFacilityLocation(ctx context.Context, arg *CreateFacilityLocationParams) (*FacilityLocation, error) {
	row := q.db.QueryRowContext(ctx, createFacilityLocation, arg.Name, arg.FacilityID)
	var i FacilityLocation
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.FacilityID,
	)
	return &i, err
}

const deleteFacilityLocation = `-- name: DeleteFacilityLocation :exec
DELETE
FROM facility_locations
WHERE id = $1
`

// description: Delete a location within a facility
func (q *Queries) DeleteFacilityLocation(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteFacilityLocation, id)
	return err
}

const getFacilityLocationByID = `-- name: GetFacilityLocationByID :one
SELECT id, created_at, updated_at, name, facility_id
FROM facility_locations
WHERE id = $1
LIMIT 1
`

// description: Get a location within a facility by ID
func (q *Queries) GetFacilityLocationByID(ctx context.Context, id int64) (*FacilityLocation, error) {
	row := q.db.QueryRowContext(ctx, getFacilityLocationByID, id)
	var i FacilityLocation
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.FacilityID,
	)
	return &i, err
}

const getFacilityLocationByName = `-- name: GetFacilityLocationByName :one
SELECT id, created_at, updated_at, name, facility_id
FROM facility_locations
WHERE name ILIKE $1
LIMIT 1
`

// description: Get a location within a facility by name
func (q *Queries) GetFacilityLocationByName(ctx context.Context, name string) (*FacilityLocation, error) {
	row := q.db.QueryRowContext(ctx, getFacilityLocationByName, name)
	var i FacilityLocation
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.FacilityID,
	)
	return &i, err
}

const listFacilityLocations = `-- name: ListFacilityLocations :many
SELECT id, created_at, updated_at, name, facility_id
FROM facility_locations
`

// description: List all locations within facilities
func (q *Queries) ListFacilityLocations(ctx context.Context) ([]*FacilityLocation, error) {
	rows, err := q.db.QueryContext(ctx, listFacilityLocations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*FacilityLocation{}
	for rows.Next() {
		var i FacilityLocation
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.FacilityID,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listFacilityLocationsByFacility = `-- name: ListFacilityLocationsByFacility :many
SELECT id, created_at, updated_at, name, facility_id
FROM facility_locations
WHERE facility_id = $1
`

// description: List all locations within a facility
func (q *Queries) ListFacilityLocationsByFacility(ctx context.Context, facilityID int64) ([]*FacilityLocation, error) {
	rows, err := q.db.QueryContext(ctx, listFacilityLocationsByFacility, facilityID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*FacilityLocation{}
	for rows.Next() {
		var i FacilityLocation
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Name,
			&i.FacilityID,
		); err != nil {
			return nil, err
		}
		items = append(items, &i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateFacilityLocation = `-- name: UpdateFacilityLocation :one
UPDATE facility_locations
SET name        = COALESCE($1, name),
    facility_id = COALESCE($2, facility_id),
    updated_at  = NOW()
WHERE id = $3
RETURNING id, created_at, updated_at, name, facility_id
`

type UpdateFacilityLocationParams struct {
	Name       nulls.String `json:"name"`
	FacilityID nulls.Int64  `json:"facility_id"`
	ID         int64        `json:"id"`
}

// description: Update a location within a facility
func (q *Queries) UpdateFacilityLocation(ctx context.Context, arg *UpdateFacilityLocationParams) (*FacilityLocation, error) {
	row := q.db.QueryRowContext(ctx, updateFacilityLocation, arg.Name, arg.FacilityID, arg.ID)
	var i FacilityLocation
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.FacilityID,
	)
	return &i, err
}
