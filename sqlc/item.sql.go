// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: item.sql

package sqlc

import (
	"context"
	"time"

	"github.com/gobuffalo/nulls"
)

const createItem = `-- name: CreateItem :one
INSERT INTO items (description, is_used, item_type_id, strain_id)
VALUES ($1, $2, $3, $4)
RETURNING id, created_at, updated_at, description, is_used, item_type_id, strain_id
`

type CreateItemParams struct {
	Description string `json:"description"`
	IsUsed      bool   `json:"is_used"`
	ItemTypeID  int64  `json:"item_type_id"`
	StrainID    int64  `json:"strain_id"`
}

// description: Create a new item
func (q *Queries) CreateItem(ctx context.Context, arg *CreateItemParams) (*Item, error) {
	row := q.db.QueryRowContext(ctx, createItem,
		arg.Description,
		arg.IsUsed,
		arg.ItemTypeID,
		arg.StrainID,
	)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Description,
		&i.IsUsed,
		&i.ItemTypeID,
		&i.StrainID,
	)
	return &i, err
}

const deleteItem = `-- name: DeleteItem :exec
DELETE
FROM items
WHERE id = $1
`

// description: Delete an item by ID
func (q *Queries) DeleteItem(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteItem, id)
	return err
}

const getItem = `-- name: GetItem :one
SELECT id, created_at, updated_at, description, is_used, item_type_id, strain_id
FROM items
WHERE id = $1
LIMIT 1
`

// description: Get an item by ID
func (q *Queries) GetItem(ctx context.Context, id int64) (*Item, error) {
	row := q.db.QueryRowContext(ctx, getItem, id)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Description,
		&i.IsUsed,
		&i.ItemTypeID,
		&i.StrainID,
	)
	return &i, err
}

const listItems = `-- name: ListItems :many
SELECT i.id, i.created_at, i.updated_at, i.description, i.is_used, i.item_type_id, i.strain_id, s.name AS strain_name, t.product_form AS product_form, t.product_modifier AS product_modifier
FROM items i
         INNER JOIN strains s ON i.strain_id = s.id
         INNER JOIN item_types t ON i.item_type_id = t.id
ORDER BY strain_id
`

type ListItemsRow struct {
	ID              int64     `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	Description     string    `json:"description"`
	IsUsed          bool      `json:"is_used"`
	ItemTypeID      int64     `json:"item_type_id"`
	StrainID        int64     `json:"strain_id"`
	StrainName      string    `json:"strain_name"`
	ProductForm     string    `json:"product_form"`
	ProductModifier string    `json:"product_modifier"`
}

// description: List all items
func (q *Queries) ListItems(ctx context.Context) ([]*ListItemsRow, error) {
	rows, err := q.db.QueryContext(ctx, listItems)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*ListItemsRow{}
	for rows.Next() {
		var i ListItemsRow
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Description,
			&i.IsUsed,
			&i.ItemTypeID,
			&i.StrainID,
			&i.StrainName,
			&i.ProductForm,
			&i.ProductModifier,
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

const updateItem = `-- name: UpdateItem :one
UPDATE items
SET description  = COALESCE($1, description),
    is_used      = COALESCE($2, is_used),
    item_type_id = COALESCE($3, item_type_id),
    strain_id    = COALESCE($4, strain_id)
WHERE id = $5
RETURNING id, created_at, updated_at, description, is_used, item_type_id, strain_id
`

type UpdateItemParams struct {
	Description nulls.String `json:"description"`
	IsUsed      nulls.Bool   `json:"is_used"`
	ItemTypeID  nulls.Int64  `json:"item_type_id"`
	StrainID    nulls.Int64  `json:"strain_id"`
	ID          int64        `json:"id"`
}

// description: Update an item by ID
func (q *Queries) UpdateItem(ctx context.Context, arg *UpdateItemParams) (*Item, error) {
	row := q.db.QueryRowContext(ctx, updateItem,
		arg.Description,
		arg.IsUsed,
		arg.ItemTypeID,
		arg.StrainID,
		arg.ID,
	)
	var i Item
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Description,
		&i.IsUsed,
		&i.ItemTypeID,
		&i.StrainID,
	)
	return &i, err
}
