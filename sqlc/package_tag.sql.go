// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: package_tag.sql

package sqlc

import (
	"context"

	"github.com/gobuffalo/nulls"
)

const createPackageTag = `-- name: CreatePackageTag :one
INSERT INTO package_tags (tag_number, is_assigned, is_provisional, is_active, assigned_package_id)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, created_at, updated_at, tag_number, is_assigned, is_provisional, is_active, assigned_package_id
`

type CreatePackageTagParams struct {
	TagNumber         string      `json:"tag_number"`
	IsAssigned        bool        `json:"is_assigned"`
	IsProvisional     bool        `json:"is_provisional"`
	IsActive          bool        `json:"is_active"`
	AssignedPackageID nulls.Int64 `json:"assigned_package_id"`
}

// description: Create a package tag
func (q *Queries) CreatePackageTag(ctx context.Context, arg *CreatePackageTagParams) (*PackageTag, error) {
	row := q.db.QueryRowContext(ctx, createPackageTag,
		arg.TagNumber,
		arg.IsAssigned,
		arg.IsProvisional,
		arg.IsActive,
		arg.AssignedPackageID,
	)
	var i PackageTag
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TagNumber,
		&i.IsAssigned,
		&i.IsProvisional,
		&i.IsActive,
		&i.AssignedPackageID,
	)
	return &i, err
}

const getPackageTagByID = `-- name: GetPackageTagByID :one
SELECT id, created_at, updated_at, tag_number, is_assigned, is_provisional, is_active, assigned_package_id
FROM package_tags
WHERE id = $1
LIMIT 1
`

// description: Get a package tag by ID
func (q *Queries) GetPackageTagByID(ctx context.Context, id int64) (*PackageTag, error) {
	row := q.db.QueryRowContext(ctx, getPackageTagByID, id)
	var i PackageTag
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TagNumber,
		&i.IsAssigned,
		&i.IsProvisional,
		&i.IsActive,
		&i.AssignedPackageID,
	)
	return &i, err
}

const getPackageTagByTagNumber = `-- name: GetPackageTagByTagNumber :one
SELECT id, created_at, updated_at, tag_number, is_assigned, is_provisional, is_active, assigned_package_id
FROM package_tags
WHERE tag_number ILIKE $1
LIMIT 1
`

// description: Get a package tag by tag number
func (q *Queries) GetPackageTagByTagNumber(ctx context.Context, tagNumber string) (*PackageTag, error) {
	row := q.db.QueryRowContext(ctx, getPackageTagByTagNumber, tagNumber)
	var i PackageTag
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TagNumber,
		&i.IsAssigned,
		&i.IsProvisional,
		&i.IsActive,
		&i.AssignedPackageID,
	)
	return &i, err
}

const listPackageTags = `-- name: ListPackageTags :many
SELECT id, created_at, updated_at, tag_number, is_assigned, is_provisional, is_active, assigned_package_id
FROM package_tags
WHERE is_assigned = $1
ORDER BY id
LIMIT $2 OFFSET $3
`

type ListPackageTagsParams struct {
	IsAssigned bool  `json:"is_assigned"`
	Limit      int32 `json:"limit"`
	Offset     int32 `json:"offset"`
}

// description: List package tags by limit and offset
func (q *Queries) ListPackageTags(ctx context.Context, arg *ListPackageTagsParams) ([]*PackageTag, error) {
	rows, err := q.db.QueryContext(ctx, listPackageTags, arg.IsAssigned, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*PackageTag{}
	for rows.Next() {
		var i PackageTag
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.TagNumber,
			&i.IsAssigned,
			&i.IsProvisional,
			&i.IsActive,
			&i.AssignedPackageID,
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

const updatePackageTag = `-- name: UpdatePackageTag :one
UPDATE package_tags
SET is_assigned         = COALESCE($1, is_assigned),
    is_provisional      = COALESCE($2, is_provisional),
    is_active           = COALESCE($3, is_active),
    assigned_package_id = COALESCE($4, assigned_package_id),
    updated_at          = NOW()
WHERE id = $5
RETURNING id, created_at, updated_at, tag_number, is_assigned, is_provisional, is_active, assigned_package_id
`

type UpdatePackageTagParams struct {
	IsAssigned        nulls.Bool  `json:"is_assigned"`
	IsProvisional     nulls.Bool  `json:"is_provisional"`
	IsActive          nulls.Bool  `json:"is_active"`
	AssignedPackageID nulls.Int64 `json:"assigned_package_id"`
	ID                int64       `json:"id"`
}

// description: Update a package tag
func (q *Queries) UpdatePackageTag(ctx context.Context, arg *UpdatePackageTagParams) (*PackageTag, error) {
	row := q.db.QueryRowContext(ctx, updatePackageTag,
		arg.IsAssigned,
		arg.IsProvisional,
		arg.IsActive,
		arg.AssignedPackageID,
		arg.ID,
	)
	var i PackageTag
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.TagNumber,
		&i.IsAssigned,
		&i.IsProvisional,
		&i.IsActive,
		&i.AssignedPackageID,
	)
	return &i, err
}
