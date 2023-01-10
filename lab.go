package pxthc

import (
	"context"
	"encoding/json"
	"github.com/earlofurl/pxthc/sqlc"
	"github.com/shopspring/decimal"
	"io"
	"time"
)

type LabTestService interface {
	FindAllLabTests(ctx context.Context) ([]*sqlc.LabTest, error)
	FindLabTestByID(ctx context.Context, id int64) (*sqlc.LabTest, error)
	CreateLabTest(ctx context.Context, l *sqlc.CreateLabTestParams) (*sqlc.LabTest, error)
	UpdateLabTest(ctx context.Context, l *sqlc.UpdateLabTestParams) (*sqlc.LabTest, error)
	DeleteLabTest(ctx context.Context, id int64) error
}

type CreateLabTestRequest struct {
	TestName                string          `json:"test_name" binding:"required"`
	BatchCode               string          `json:"batch_code" binding:"required"`
	TestIDCode              string          `json:"test_id_code" binding:"required"`
	LabFacilityName         string          `json:"lab_facility_name" binding:"required"`
	TestPerformedDateTime   time.Time       `json:"test_performed_date_time" binding:"required"`
	OverallPassed           bool            `json:"overall_passed" binding:"required"`
	TestTypeName            string          `json:"test_type_name" binding:"required"`
	TestPassed              bool            `json:"test_passed" binding:"required"`
	TestComment             string          `json:"test_comment" binding:"required"`
	ThcTotalPercent         decimal.Decimal `json:"thc_total_percent" binding:"required"`
	ThcTotalValue           decimal.Decimal `json:"thc_total_value" binding:"required"`
	CbdPercent              decimal.Decimal `json:"cbd_percent" binding:"required"`
	CbdValue                decimal.Decimal `json:"cbd_value" binding:"required"`
	TerpeneTotalPercent     decimal.Decimal `json:"terpene_total_percent" binding:"required"`
	TerpeneTotalValue       decimal.Decimal `json:"terpene_total_value" binding:"required"`
	ThcAPercent             decimal.Decimal `json:"thc_a_percent" binding:"required"`
	ThcAValue               decimal.Decimal `json:"thc_a_value" binding:"required"`
	Delta9ThcPercent        decimal.Decimal `json:"delta_9_thc_percent" binding:"required"`
	Delta9ThcValue          decimal.Decimal `json:"delta_9_thc_value" binding:"required"`
	Delta8ThcPercent        decimal.Decimal `json:"delta_8_thc_percent" binding:"required"`
	Delta8ThcValue          decimal.Decimal `json:"delta_8_thc_value" binding:"required"`
	ThcVPercent             decimal.Decimal `json:"thc_v_percent" binding:"required"`
	ThcVValue               decimal.Decimal `json:"thc_v_value" binding:"required"`
	CbdAPercent             decimal.Decimal `json:"cbd_a_percent" binding:"required"`
	CbdAValue               decimal.Decimal `json:"cbd_a_value" binding:"required"`
	CbnPercent              decimal.Decimal `json:"cbn_percent" binding:"required"`
	CbnValue                decimal.Decimal `json:"cbn_value" binding:"required"`
	CbgAPercent             decimal.Decimal `json:"cbg_a_percent" binding:"required"`
	CbgAValue               decimal.Decimal `json:"cbg_a_value" binding:"required"`
	CbgPercent              decimal.Decimal `json:"cbg_percent" binding:"required"`
	CbgValue                decimal.Decimal `json:"cbg_value" binding:"required"`
	CbcPercent              decimal.Decimal `json:"cbc_percent" binding:"required"`
	CbcValue                decimal.Decimal `json:"cbc_value" binding:"required"`
	TotalCannabinoidPercent decimal.Decimal `json:"total_cannabinoid_percent" binding:"required"`
	TotalCannabinoidValue   decimal.Decimal `json:"total_cannabinoid_value" binding:"required"`
}

type UpdateLabTestRequest struct {
	TestName                string          `json:"test_name" binding:"required"`
	BatchCode               string          `json:"batch_code" binding:"required"`
	TestIDCode              string          `json:"test_id_code" binding:"required"`
	LabFacilityName         string          `json:"lab_facility_name" binding:"required"`
	TestPerformedDateTime   time.Time       `json:"test_performed_date_time" binding:"required"`
	OverallPassed           bool            `json:"overall_passed" binding:"required"`
	TestTypeName            string          `json:"test_type_name" binding:"required"`
	TestPassed              bool            `json:"test_passed" binding:"required"`
	TestComment             string          `json:"test_comment" binding:"required"`
	ThcTotalPercent         decimal.Decimal `json:"thc_total_percent" binding:"required"`
	ThcTotalValue           decimal.Decimal `json:"thc_total_value" binding:"required"`
	CbdPercent              decimal.Decimal `json:"cbd_percent" binding:"required"`
	CbdValue                decimal.Decimal `json:"cbd_value" binding:"required"`
	TerpeneTotalPercent     decimal.Decimal `json:"terpene_total_percent" binding:"required"`
	TerpeneTotalValue       decimal.Decimal `json:"terpene_total_value" binding:"required"`
	ThcAPercent             decimal.Decimal `json:"thc_a_percent" binding:"required"`
	ThcAValue               decimal.Decimal `json:"thc_a_value" binding:"required"`
	Delta9ThcPercent        decimal.Decimal `json:"delta_9_thc_percent" binding:"required"`
	Delta9ThcValue          decimal.Decimal `json:"delta_9_thc_value" binding:"required"`
	Delta8ThcPercent        decimal.Decimal `json:"delta_8_thc_percent" binding:"required"`
	Delta8ThcValue          decimal.Decimal `json:"delta_8_thc_value" binding:"required"`
	ThcVPercent             decimal.Decimal `json:"thc_v_percent" binding:"required"`
	ThcVValue               decimal.Decimal `json:"thc_v_value" binding:"required"`
	CbdAPercent             decimal.Decimal `json:"cbd_a_percent" binding:"required"`
	CbdAValue               decimal.Decimal `json:"cbd_a_value" binding:"required"`
	CbnPercent              decimal.Decimal `json:"cbn_percent" binding:"required"`
	CbnValue                decimal.Decimal `json:"cbn_value" binding:"required"`
	CbgAPercent             decimal.Decimal `json:"cbg_a_percent" binding:"required"`
	CbgAValue               decimal.Decimal `json:"cbg_a_value" binding:"required"`
	CbgPercent              decimal.Decimal `json:"cbg_percent" binding:"required"`
	CbgValue                decimal.Decimal `json:"cbg_value" binding:"required"`
	CbcPercent              decimal.Decimal `json:"cbc_percent" binding:"required"`
	CbcValue                decimal.Decimal `json:"cbc_value" binding:"required"`
	TotalCannabinoidPercent decimal.Decimal `json:"total_cannabinoid_percent" binding:"required"`
	TotalCannabinoidValue   decimal.Decimal `json:"total_cannabinoid_value" binding:"required"`
	ID                      int64           `json:"id" binding:"required"`
}

func (r *CreateLabTestRequest) Bind(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(r)
}

func (r *UpdateLabTestRequest) Bind(body io.ReadCloser) error {
	return json.NewDecoder(body).Decode(r)
}
