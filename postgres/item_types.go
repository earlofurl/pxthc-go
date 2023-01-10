package postgres

import (
	"context"
	"github.com/earlofurl/pxthc"
	"github.com/earlofurl/pxthc/sqlc"
)

// Ensure service implements interface.
var _ pxthc.ItemTypeService = (*ItemTypeService)(nil)

type ItemTypeService struct {
	store sqlc.Store
}

func NewItemTypeService(store *sqlc.Store) *ItemTypeService {
	return &ItemTypeService{store: *store}
}

// FindAllItemTypes retrieves all item types.
func (s *ItemTypeService) FindAllItemTypes(ctx context.Context) ([]*sqlc.ItemType, error) {
	u, err := s.store.ListItemTypes(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// FindItemTypeByID retrieves an item type by ID.
// Returns ENOTFOUND if item type does not exist.
func (s *ItemTypeService) FindItemTypeByID(ctx context.Context, id int64) (*sqlc.ItemType, error) {
	u, err := s.store.GetItemTypeByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, pxthc.ErrNoRecord
	}
	return u, nil
}

// CreateItemType creates a new item type.
func (s *ItemTypeService) CreateItemType(ctx context.Context, arg *sqlc.CreateItemTypeParams) (*sqlc.ItemType, error) {
	u, err := s.store.CreateItemType(ctx, arg)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// UpdateItemType updates an item type.
func (s *ItemTypeService) UpdateItemType(ctx context.Context, arg *sqlc.UpdateItemTypeParams) (*sqlc.ItemType, error) {
	u, err := s.store.UpdateItemType(ctx, arg)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// DeleteItemType deletes an item type.
func (s *ItemTypeService) DeleteItemType(ctx context.Context, id int64) error {
	return s.store.DeleteItemType(ctx, id)
}
