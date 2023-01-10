package postgres

import (
	"context"
	"github.com/earlofurl/pxthc"
	"github.com/earlofurl/pxthc/sqlc"
	"github.com/gobuffalo/nulls"
)

// ensure service implements interface
var _ pxthc.LabTestService = (*LabTestService)(nil)

type LabTestService struct {
	store sqlc.Store
}

func NewLabTestService(store *sqlc.Store) *LabTestService {
	return &LabTestService{store: *store}
}

// FindAllLabTests retrieves all lab tests.
func (s *LabTestService) FindAllLabTests(ctx context.Context) ([]*sqlc.LabTest, error) {
	u, err := s.store.ListLabTests(ctx)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// FindLabTestByID retrieves a lab test by ID.
// Returns ENOTFOUND if lab test does not exist.
func (s *LabTestService) FindLabTestByID(ctx context.Context, id int64) (*sqlc.LabTest, error) {
	u, err := s.store.GetLabTestByID(ctx, nulls.NewInt64(id))
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, pxthc.ErrNoRecord
	}
	return u, nil
}

// CreateLabTest creates a new lab test.
func (s *LabTestService) CreateLabTest(ctx context.Context, f *sqlc.CreateLabTestParams) (*sqlc.LabTest, error) {
	u, err := s.store.CreateLabTest(ctx, f)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// UpdateLabTest updates a lab test.
func (s *LabTestService) UpdateLabTest(ctx context.Context, p *sqlc.UpdateLabTestParams) (*sqlc.LabTest, error) {
	u, err := s.store.UpdateLabTest(ctx, p)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// DeleteLabTest deletes a lab test.
func (s *LabTestService) DeleteLabTest(ctx context.Context, id int64) error {
	err := s.store.DeleteLabTest(ctx, nulls.NewInt64(id))
	if err != nil {
		return err
	}
	return nil
}
