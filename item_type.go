package pxthc

import (
	"context"
	"encoding/json"
	"github.com/earlofurl/pxthc/sqlc"
	"io"
)

type ItemTypeService interface {
	FindAllItemTypes(ctx context.Context) ([]*sqlc.ItemType, error)
	FindItemTypeByID(ctx context.Context, id int64) (*sqlc.ItemType, error)
	CreateItemType(ctx context.Context, f *sqlc.CreateItemTypeParams) (*sqlc.ItemType, error)
	UpdateItemType(ctx context.Context, f *sqlc.UpdateItemTypeParams) (*sqlc.ItemType, error)
	DeleteItemType(ctx context.Context, id int64) error
}

type CreateItemTypeRequest struct {
	ProductForm       string `json:"product_form"`
	ProductModifier   string `json:"product_modifier"`
	UomDefault        int64  `json:"uom_default"`
	ProductCategoryID int64  `json:"product_category_id"`
}

type UpdateItemTypeRequest struct {
	ProductForm       string `json:"product_form"`
	ProductModifier   string `json:"product_modifier"`
	UomDefault        int64  `json:"uom_default"`
	ProductCategoryID int64  `json:"product_category_id"`
	ID                int64  `json:"id"`
}

func (r *CreateItemTypeRequest) Bind(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *UpdateItemTypeRequest) Bind(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(r)
}
