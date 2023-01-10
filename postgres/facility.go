package postgres

import (
	"context"
	"github.com/earlofurl/pxthc"
	"github.com/earlofurl/pxthc/sqlc"
)

// Ensure service implements interface.
var _ pxthc.FacilityService = (*FacilityService)(nil)

type FacilityService struct {
	store sqlc.Store
}

func NewFacilityService(store *sqlc.Store) *FacilityService {
	return &FacilityService{store: *store}
}

// FindAllFacilities retrieves all facilities.
func (s *FacilityService) FindAllFacilities(ctx context.Context) ([]*sqlc.Facility, error) {
	u, err := s.store.ListFacilities(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// FindFacilityByID retrieves a facility by ID.
// Returns ENOTFOUND if facility does not exist.
func (s *FacilityService) FindFacilityByID(ctx context.Context, id int64) (*sqlc.Facility, error) {
	u, err := s.store.GetFacilityByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, pxthc.ErrNoRecord
	}
	return u, nil
}

// FindFacilityByName retrieves a facility by Name.
// Returns ENOTFOUND if facility does not exist.
func (s *FacilityService) FindFacilityByName(ctx context.Context, name string) (*sqlc.Facility, error) {
	u, err := s.store.GetFacilityByName(ctx, name)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, pxthc.ErrNoRecord
	}
	return u, nil
}

// CreateFacility creates a new facility.
func (s *FacilityService) CreateFacility(ctx context.Context, f *sqlc.CreateFacilityParams) (*sqlc.Facility, error) {
	u, err := s.store.CreateFacility(ctx, f)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// UpdateFacility updates a facility.
func (s *FacilityService) UpdateFacility(ctx context.Context, f *sqlc.UpdateFacilityParams) (*sqlc.Facility, error) {
	u, err := s.store.UpdateFacility(ctx, f)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// DeleteFacility deletes a facility.
func (s *FacilityService) DeleteFacility(ctx context.Context, id int64) error {
	err := s.store.DeleteFacility(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
