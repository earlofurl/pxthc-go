package pxthc

import (
	"context"
	"encoding/json"
	"github.com/earlofurl/pxthc/sqlc"
	"io"
)

type FacilityLocationService interface {
	FindAllFacilityLocations(ctx context.Context) ([]*sqlc.FacilityLocation, error)
	FindFacilityLocationByID(ctx context.Context, id int64) (*sqlc.FacilityLocation, error)
	FindFacilityLocationByName(ctx context.Context, name string) (*sqlc.FacilityLocation, error)
	CreateFacilityLocation(ctx context.Context, f *sqlc.CreateFacilityLocationParams) (*sqlc.FacilityLocation, error)
	UpdateFacilityLocation(ctx context.Context, f *sqlc.UpdateFacilityLocationParams) (*sqlc.FacilityLocation, error)
	DeleteFacilityLocation(ctx context.Context, id int64) error
}

type CreateFacilityLocationRequest struct {
	Name       string `json:"name"`
	FacilityID int64  `json:"facility_id"`
}

type UpdateFacilityLocationRequest struct {
	Name       string `json:"name"`
	FacilityID int64  `json:"facility_id"`
	ID         int64  `json:"id"`
}

func (r *CreateFacilityLocationRequest) Bind(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *UpdateFacilityLocationRequest) Bind(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(r)
}
