package pxthc

import (
	"context"
	"github.com/earlofurl/pxthc/sqlc"
	"github.com/gobuffalo/nulls"
	"github.com/shopspring/decimal"
)

type StrainService interface {
	FindAllStrains(ctx context.Context) ([]*sqlc.Strain, error)
	FindStrainByID(ctx context.Context, id int64) (*sqlc.Strain, error)
	FindStrainByName(ctx context.Context, name string) (*sqlc.Strain, error)
	UpdateStrain(ctx context.Context, r *sqlc.UpdateStrainParams) (*sqlc.Strain, error)
	CreateStrain(ctx context.Context, r *sqlc.CreateStrainParams) (*sqlc.Strain, error)
	DeleteStrain(ctx context.Context, id int64) error
}

type CreateStrainRequest struct {
	Name string `json:"name" binding:"required"`
	Type string `json:"type" binding:"required"`
}

type updateStrainRequest struct {
	ID                      int64               `json:"id" binding:"required,min=1"`
	Name                    nulls.String        `json:"name"`
	Type                    nulls.String        `json:"type"`
	YieldAverage            decimal.NullDecimal `json:"yield_average"`
	TerpAverageTotal        decimal.NullDecimal `json:"terp_average_total"`
	Terp1                   nulls.String        `json:"terp_1"`
	Terp1Value              decimal.NullDecimal `json:"terp_1_value"`
	Terp2                   nulls.String        `json:"terp_2"`
	Terp2Value              decimal.NullDecimal `json:"terp_2_value"`
	Terp3                   nulls.String        `json:"terp_3"`
	Terp3Value              decimal.NullDecimal `json:"terp_3_value"`
	Terp4                   nulls.String        `json:"terp_4"`
	Terp4Value              decimal.NullDecimal `json:"terp_4_value"`
	Terp5                   nulls.String        `json:"terp_5"`
	Terp5Value              decimal.NullDecimal `json:"terp_5_value"`
	ThcAverage              decimal.NullDecimal `json:"thc_average"`
	TotalCannabinoidAverage decimal.NullDecimal `json:"total_cannabinoid_average"`
	LightDep2022            nulls.String        `json:"light_dep_2022"`
	FallHarvest2022         nulls.String        `json:"fall_harvest_2022"`
	QuantityAvailable       decimal.NullDecimal `json:"quantity_available"`
}
