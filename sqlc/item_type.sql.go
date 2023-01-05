// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: item_type.sql

package sqlc

import (
	"context"

	"github.com/gobuffalo/nulls"
)

const createItemType = `-- name: CreateItemType :one
INSERT INTO item_types (product_form,
                        product_modifier,
                        uom_default,
                        product_category_id)
VALUES ($1, $2, $3, $4)
RETURNING id, created_at, updated_at, product_form, product_modifier, uom_default, product_category_id
`

type CreateItemTypeParams struct {
	ProductForm       string `json:"product_form"`
	ProductModifier   string `json:"product_modifier"`
	UomDefault        int64  `json:"uom_default"`
	ProductCategoryID int64  `json:"product_category_id"`
}

// description: Create a new item type
func (q *Queries) CreateItemType(ctx context.Context, arg *CreateItemTypeParams) (*ItemType, error) {
	row := q.db.QueryRowContext(ctx, createItemType,
		arg.ProductForm,
		arg.ProductModifier,
		arg.UomDefault,
		arg.ProductCategoryID,
	)
	var i ItemType
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProductForm,
		&i.ProductModifier,
		&i.UomDefault,
		&i.ProductCategoryID,
	)
	return &i, err
}

const deleteItemType = `-- name: DeleteItemType :exec
DELETE
FROM item_types
WHERE id = $1
`

// description: Delete an item type by ID
func (q *Queries) DeleteItemType(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteItemType, id)
	return err
}

const getItemType = `-- name: GetItemType :one
SELECT id, created_at, updated_at, product_form, product_modifier, uom_default, product_category_id
FROM item_types
WHERE id = $1
LIMIT 1
`

// description: Get an item type by id
func (q *Queries) GetItemType(ctx context.Context, id int64) (*ItemType, error) {
	row := q.db.QueryRowContext(ctx, getItemType, id)
	var i ItemType
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProductForm,
		&i.ProductModifier,
		&i.UomDefault,
		&i.ProductCategoryID,
	)
	return &i, err
}

const listItemTypes = `-- name: ListItemTypes :many
SELECT id, created_at, updated_at, product_form, product_modifier, uom_default, product_category_id
FROM item_types
ORDER BY id
`

// description: List all item types
func (q *Queries) ListItemTypes(ctx context.Context) ([]*ItemType, error) {
	rows, err := q.db.QueryContext(ctx, listItemTypes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []*ItemType{}
	for rows.Next() {
		var i ItemType
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.ProductForm,
			&i.ProductModifier,
			&i.UomDefault,
			&i.ProductCategoryID,
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

const updateItemType = `-- name: UpdateItemType :one
UPDATE item_types
SET product_form        = COALESCE($1, product_form),
    product_modifier    = COALESCE($2, product_modifier),
    uom_default         = COALESCE($3, uom_default),
    product_category_id = COALESCE($4, product_category_id),
    updated_at          = NOW()
WHERE id = $5
    RETURNING id, created_at, updated_at, product_form, product_modifier, uom_default, product_category_id
`

type UpdateItemTypeParams struct {
	ProductForm       nulls.String `json:"product_form"`
	ProductModifier   nulls.String `json:"product_modifier"`
	UomDefault        nulls.Int64  `json:"uom_default"`
	ProductCategoryID nulls.Int64  `json:"product_category_id"`
	ID                int64        `json:"id"`
}

// description: Update an item type by ID
func (q *Queries) UpdateItemType(ctx context.Context, arg *UpdateItemTypeParams) (*ItemType, error) {
	row := q.db.QueryRowContext(ctx, updateItemType,
		arg.ProductForm,
		arg.ProductModifier,
		arg.UomDefault,
		arg.ProductCategoryID,
		arg.ID,
	)
	var i ItemType
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ProductForm,
		&i.ProductModifier,
		&i.UomDefault,
		&i.ProductCategoryID,
	)
	return &i, err
}