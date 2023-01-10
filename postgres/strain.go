package postgres

import (
	"context"
	"github.com/earlofurl/pxthc"
	"github.com/earlofurl/pxthc/sqlc"
)

// Ensure service implements interface.
var _ pxthc.StrainService = (*StrainService)(nil)

type StrainService struct {
	store sqlc.Store
}

func NewStrainService(store *sqlc.Store) *StrainService {
	return &StrainService{store: *store}
}

// CreateStrain creates a new strain.
func (s *StrainService) CreateStrain(ctx context.Context, arg *sqlc.CreateStrainParams) (*sqlc.Strain, error) {
	strain, err := s.store.CreateStrain(ctx, arg)
	if err != nil {
		return nil, err
	}
	return strain, nil
}

// FindAllStrains retrieves all strains.
func (s *StrainService) FindAllStrains(ctx context.Context) ([]*sqlc.Strain, error) {
	strains, err := s.store.ListStrains(ctx)
	if err != nil {
		return nil, err
	}
	return strains, nil
}

// FindStrainByID retrieves a strain by ID.
// Returns ENOTFOUND if strain does not exist.
func (s *StrainService) FindStrainByID(ctx context.Context, id int64) (*sqlc.Strain, error) {
	strain, err := s.store.GetStrainByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if strain == nil {
		return nil, pxthc.ErrNoRecord
	}
	return strain, nil
}

// FindStrainByName retrieves a strain by Name.
// Returns ENOTFOUND if strain does not exist.
func (s *StrainService) FindStrainByName(ctx context.Context, name string) (*sqlc.Strain, error) {
	strain, err := s.store.GetStrainByName(ctx, name)
	if err != nil {
		return nil, err
	}
	if strain == nil {
		return nil, pxthc.ErrNoRecord
	}
	return strain, nil
}

// UpdateStrain updates a strain.
func (s *StrainService) UpdateStrain(ctx context.Context, arg *sqlc.UpdateStrainParams) (*sqlc.Strain, error) {
	strain, err := s.store.UpdateStrain(ctx, arg)
	if err != nil {
		return nil, err
	}
	return strain, nil
}

// DeleteStrain deletes a strain.
func (s *StrainService) DeleteStrain(ctx context.Context, id int64) error {
	err := s.store.DeleteStrain(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
