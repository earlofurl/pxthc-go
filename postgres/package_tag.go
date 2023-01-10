package postgres

import (
	"context"
	"github.com/earlofurl/pxthc"
	"github.com/earlofurl/pxthc/sqlc"
)

// Ensure service implements interface.
var _ pxthc.PackageTagService = (*PackageTagService)(nil)

type PackageTagService struct {
	store sqlc.Store
}

func NewPackageTagService(store *sqlc.Store) *PackageTagService {
	return &PackageTagService{store: *store}
}

// FindAllPackageTags retrieves all package tags.
func (s *PackageTagService) FindAllPackageTags(ctx context.Context, arg *sqlc.ListPackageTagsParams) ([]*sqlc.PackageTag, error) {
	u, err := s.store.ListPackageTags(ctx, arg)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// FindPackageTagByID retrieves a package tag by ID.
// Returns ENOTFOUND if package tag does not exist.
func (s *PackageTagService) FindPackageTagByID(ctx context.Context, id int64) (*sqlc.PackageTag, error) {
	u, err := s.store.GetPackageTagByID(ctx, id)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, pxthc.ErrNoRecord
	}
	return u, nil
}

// FindPackageTagByTagNumber retrieves a package tag by TagNumber.
// Returns ENOTFOUND if package tag does not exist.
func (s *PackageTagService) FindPackageTagByTagNumber(ctx context.Context, tagNumber string) (*sqlc.PackageTag, error) {
	u, err := s.store.GetPackageTagByTagNumber(ctx, tagNumber)
	if err != nil {
		return nil, err
	}
	if u == nil {
		return nil, pxthc.ErrNoRecord
	}
	return u, nil
}

// UpdatePackageTag updates a package tag.
func (s *PackageTagService) UpdatePackageTag(ctx context.Context, p *sqlc.UpdatePackageTagParams) (*sqlc.PackageTag, error) {
	u, err := s.store.UpdatePackageTag(ctx, p)
	if err != nil {
		return nil, err
	}
	return u, nil
}
