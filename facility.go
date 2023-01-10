package pxthc

import (
	"context"
	"encoding/json"
	"github.com/earlofurl/pxthc/sqlc"
	"io"
)

type FacilityService interface {
	FindAllFacilities(ctx context.Context) ([]*sqlc.Facility, error)
	FindFacilityByID(ctx context.Context, id int64) (*sqlc.Facility, error)
	FindFacilityByName(ctx context.Context, name string) (*sqlc.Facility, error)
	CreateFacility(ctx context.Context, f *sqlc.CreateFacilityParams) (*sqlc.Facility, error)
	UpdateFacility(ctx context.Context, f *sqlc.UpdateFacilityParams) (*sqlc.Facility, error)
	DeleteFacility(ctx context.Context, id int64) error
}

type CreateFacilityRequest struct {
	Name          string `json:"name"`
	LicenseNumber string `json:"license_number"`
}

type UpdateFacilityRequest struct {
	Name          string `json:"name"`
	LicenseNumber string `json:"license_number"`
	ID            int64  `json:"id"`
}

func (r *CreateFacilityRequest) Bind(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *UpdateFacilityRequest) Bind(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(r)
}
