package postgres

import (
	"context"
	"github.com/earlofurl/pxthc"
	"github.com/earlofurl/pxthc/sqlc"
)

// Ensure service implements interface.
var _ pxthc.ProductCategoryService = (*ProductCategoryService)(nil)

type ProductCategoryService struct {
	store sqlc.Store
}

func NewProductCategoryService(store *sqlc.Store) *ProductCategoryService {
	return &ProductCategoryService{store: *store}
}

// FindAllProductCategories retrieves all product categories.
func (s *ProductCategoryService) FindAllProductCategories(ctx context.Context) ([]*sqlc.ProductCategory, error) {
	u, err := s.store.ListProductCategories(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// FindProductCategoryByID retrieves a product category by ID.
// Returns ENOTFOUND if product category does not exist.
func (s *ProductCategoryService) FindProductCategoryByID(ctx context.Context, id int64) (*sqlc.ProductCategory, error) {
	u, err := s.store.GetProductCategoryByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, pxthc.ErrNoRecord
	}
	return u, nil
}

// FindProductCategoryByName retrieves a product category by Name.
// Returns ENOTFOUND if product category does not exist.
func (s *ProductCategoryService) FindProductCategoryByName(ctx context.Context, name string) (*sqlc.ProductCategory, error) {
	u, err := s.store.GetProductCategoryByName(ctx, name)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, pxthc.ErrNoRecord
	}
	return u, nil
}
