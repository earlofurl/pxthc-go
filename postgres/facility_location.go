package postgres

import (
	"context"
	"github.com/earlofurl/pxthc"
	"github.com/earlofurl/pxthc/sqlc"
)

// Ensure service implements interface.
var _ pxthc.FacilityLocationService = (*FacilityLocationService)(nil)

type FacilityLocationService struct {
	store sqlc.Store
}

func NewFacilityLocationService(store *sqlc.Store) *FacilityLocationService {
	return &FacilityLocationService{store: *store}
}

// FindAllFacilityLocations retrieves all facility locations.
func (s *FacilityLocationService) FindAllFacilityLocations(ctx context.Context) ([]*sqlc.FacilityLocation, error) {
	u, err := s.store.ListFacilityLocations(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// FindFacilityLocationByID retrieves a facility location by ID.
// Returns ENOTFOUND if facility location does not exist.
func (s *FacilityLocationService) FindFacilityLocationByID(ctx context.Context, id int64) (*sqlc.FacilityLocation, error) {
	u, err := s.store.GetFacilityLocationByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, pxthc.ErrNoRecord
	}
	return u, nil
}

// FindFacilityLocationByName retrieves a facility location by Name.
// Returns ENOTFOUND if facility location does not exist.
func (s *FacilityLocationService) FindFacilityLocationByName(ctx context.Context, name string) (*sqlc.FacilityLocation, error) {
	u, err := s.store.GetFacilityLocationByName(ctx, name)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, pxthc.ErrNoRecord
	}
	return u, nil
}

// CreateFacilityLocation creates a new facility location.
func (s *FacilityLocationService) CreateFacilityLocation(ctx context.Context, f *sqlc.CreateFacilityLocationParams) (*sqlc.FacilityLocation, error) {
	u, err := s.store.CreateFacilityLocation(ctx, f)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// UpdateFacilityLocation updates a facility location.
func (s *FacilityLocationService) UpdateFacilityLocation(ctx context.Context, f *sqlc.UpdateFacilityLocationParams) (*sqlc.FacilityLocation, error) {
	u, err := s.store.UpdateFacilityLocation(ctx, f)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// DeleteFacilityLocation deletes a facility location.
func (s *FacilityLocationService) DeleteFacilityLocation(ctx context.Context, id int64) error {
	err := s.store.DeleteFacilityLocation(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
