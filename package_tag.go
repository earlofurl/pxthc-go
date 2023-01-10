package pxthc

import (
	"context"
	"encoding/json"
	"github.com/earlofurl/pxthc/sqlc"
	"io"
)

type PackageTagService interface {
	FindAllPackageTags(ctx context.Context, params *sqlc.ListPackageTagsParams) ([]*sqlc.PackageTag, error)
	FindPackageTagByID(ctx context.Context, id int64) (*sqlc.PackageTag, error)
	FindPackageTagByTagNumber(ctx context.Context, tagNumber string) (*sqlc.PackageTag, error)
	UpdatePackageTag(ctx context.Context, p *sqlc.UpdatePackageTagParams) (*sqlc.PackageTag, error)
}

type FindAllPackageTagsRequest struct {
	IsAssigned bool `json:"is_assigned"`
	Limit      int  `json:"limit"`
	Offset     int  `json:"offset"`
}

type UpdatePackageTagRequest struct {
	IsAssigned        bool  `json:"is_assigned"`
	IsProvisional     bool  `json:"is_provisional"`
	IsActive          bool  `json:"is_active"`
	AssignedPackageID int64 `json:"assigned_package_id"`
	ID                int64 `json:"id"`
}

func (r *FindAllPackageTagsRequest) Bind(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *UpdatePackageTagRequest) Bind(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(r)
}
