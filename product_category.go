package pxthc

import (
	"context"
	"github.com/earlofurl/pxthc/sqlc"
)

type ProductCategoryService interface {
	FindAllProductCategories(ctx context.Context) ([]*sqlc.ProductCategory, error)
	FindProductCategoryByID(ctx context.Context, id int64) (*sqlc.ProductCategory, error)
	FindProductCategoryByName(ctx context.Context, name string) (*sqlc.ProductCategory, error)
}
