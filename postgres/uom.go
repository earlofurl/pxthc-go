package postgres

import (
	"context"
	"github.com/earlofurl/pxthc"
	"github.com/earlofurl/pxthc/sqlc"
)

// Ensure service implements interface.
var _ pxthc.UomService = (*UomService)(nil)

type UomService struct {
	store sqlc.Store
}

func NewUomService(store *sqlc.Store) *UomService {
	return &UomService{store: *store}
}

// FindAllUoms retrieves all uoms.
func (s *UomService) FindAllUoms(ctx context.Context) ([]*sqlc.Uom, error) {
	u, err := s.store.ListUoms(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// FindUomByID retrieves a uom by ID.
// Returns ENOTFOUND if uom does not exist.
func (s *UomService) FindUomByID(ctx context.Context, id int64) (*sqlc.Uom, error) {
	u, err := s.store.GetUomByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, pxthc.ErrNoRecord
	}
	return u, nil
}

// FindUomByName retrieves a uom by Name.
// Returns ENOTFOUND if uom does not exist.
func (s *UomService) FindUomByName(ctx context.Context, name string) (*sqlc.Uom, error) {
	u, err := s.store.GetUomByName(ctx, name)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, pxthc.ErrNoRecord
	}
	return u, nil
}
