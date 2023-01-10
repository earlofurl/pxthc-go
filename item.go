package pxthc

import (
	"context"
	"encoding/json"
	"github.com/earlofurl/pxthc/sqlc"
	"io"
)

type ItemService interface {
	FindAllItems(ctx context.Context) ([]*sqlc.ListItemsRow, error)
	FindItemByID(ctx context.Context, id int64) (*sqlc.Item, error)
	CreateItem(ctx context.Context, f *sqlc.CreateItemParams) (*sqlc.Item, error)
	UpdateItem(ctx context.Context, f *sqlc.UpdateItemParams) (*sqlc.Item, error)
	DeleteItem(ctx context.Context, id int64) error
}

type CreateItemRequest struct {
	Description string `json:"description"`
	IsUsed      bool   `json:"is_used"`
	ItemTypeID  int64  `json:"item_type_id"`
	StrainID    int64  `json:"strain_id"`
}

type UpdateItemRequest struct {
	Description string `json:"description"`
	IsUsed      bool   `json:"is_used"`
	ItemTypeID  int64  `json:"item_type_id"`
	StrainID    int64  `json:"strain_id"`
	ID          int64  `json:"id"`
}

func (r *CreateItemRequest) Bind(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *UpdateItemRequest) Bind(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(r)
}
