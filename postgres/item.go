package postgres

import (
	"context"
	"github.com/earlofurl/pxthc"
	"github.com/earlofurl/pxthc/sqlc"
)

// ensure service implements interface
var _ pxthc.ItemService = (*ItemService)(nil)

type ItemService struct {
	store sqlc.Store
}

func NewItemService(store *sqlc.Store) *ItemService {
	return &ItemService{store: *store}
}

// FindAllItems retrieves all items.
func (s *ItemService) FindAllItems(ctx context.Context) ([]*sqlc.ListItemsRow, error) {
	u, err := s.store.ListItems(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// FindItemByID retrieves an item by ID.
// Returns ENOTFOUND if item does not exist.
func (s *ItemService) FindItemByID(ctx context.Context, id int64) (*sqlc.Item, error) {
	u, err := s.store.GetItemByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, pxthc.ErrNoRecord
	}
	return u, nil
}

// CreateItem creates a new item.
func (s *ItemService) CreateItem(ctx context.Context, item *sqlc.CreateItemParams) (*sqlc.Item, error) {
	u, err := s.store.CreateItem(ctx, item)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// UpdateItem updates an item.
func (s *ItemService) UpdateItem(ctx context.Context, item *sqlc.UpdateItemParams) (*sqlc.Item, error) {
	u, err := s.store.UpdateItem(ctx, item)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// DeleteItem deletes an item.
func (s *ItemService) DeleteItem(ctx context.Context, id int64) error {
	err := s.store.DeleteItem(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
