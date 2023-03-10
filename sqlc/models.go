// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package sqlc

import (
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Facility struct {
	ID            int64     `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Name          string    `json:"name"`
	LicenseNumber string    `json:"license_number"`
}

type FacilityLocation struct {
	ID         int64     `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Name       string    `json:"name"`
	FacilityID int64     `json:"facility_id"`
}

type Item struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Description string    `json:"description"`
	IsUsed      bool      `json:"is_used"`
	ItemTypeID  int64     `json:"item_type_id"`
	StrainID    int64     `json:"strain_id"`
}

type ItemType struct {
	ID                int64     `json:"id"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	ProductForm       string    `json:"product_form"`
	ProductModifier   string    `json:"product_modifier"`
	UomDefault        int64     `json:"uom_default"`
	ProductCategoryID int64     `json:"product_category_id"`
}

type LabTest struct {
	ID                      nulls.Int64     `json:"id"`
	CreatedAt               time.Time       `json:"created_at"`
	UpdatedAt               time.Time       `json:"updated_at"`
	TestName                string          `json:"test_name"`
	BatchCode               string          `json:"batch_code"`
	TestIDCode              string          `json:"test_id_code"`
	LabFacilityName         string          `json:"lab_facility_name"`
	TestPerformedDateTime   time.Time       `json:"test_performed_date_time"`
	TestCompleted           bool            `json:"test_completed"`
	OverallPassed           bool            `json:"overall_passed"`
	TestTypeName            string          `json:"test_type_name"`
	TestPassed              bool            `json:"test_passed"`
	TestComment             string          `json:"test_comment"`
	ThcTotalPercent         decimal.Decimal `json:"thc_total_percent"`
	ThcTotalValue           decimal.Decimal `json:"thc_total_value"`
	CbdPercent              decimal.Decimal `json:"cbd_percent"`
	CbdValue                decimal.Decimal `json:"cbd_value"`
	TerpeneTotalPercent     decimal.Decimal `json:"terpene_total_percent"`
	TerpeneTotalValue       decimal.Decimal `json:"terpene_total_value"`
	ThcAPercent             decimal.Decimal `json:"thc_a_percent"`
	ThcAValue               decimal.Decimal `json:"thc_a_value"`
	Delta9ThcPercent        decimal.Decimal `json:"delta9_thc_percent"`
	Delta9ThcValue          decimal.Decimal `json:"delta9_thc_value"`
	Delta8ThcPercent        decimal.Decimal `json:"delta8_thc_percent"`
	Delta8ThcValue          decimal.Decimal `json:"delta8_thc_value"`
	ThcVPercent             decimal.Decimal `json:"thc_v_percent"`
	ThcVValue               decimal.Decimal `json:"thc_v_value"`
	CbdAPercent             decimal.Decimal `json:"cbd_a_percent"`
	CbdAValue               decimal.Decimal `json:"cbd_a_value"`
	CbnPercent              decimal.Decimal `json:"cbn_percent"`
	CbnValue                decimal.Decimal `json:"cbn_value"`
	CbgAPercent             decimal.Decimal `json:"cbg_a_percent"`
	CbgAValue               decimal.Decimal `json:"cbg_a_value"`
	CbgPercent              decimal.Decimal `json:"cbg_percent"`
	CbgValue                decimal.Decimal `json:"cbg_value"`
	CbcPercent              decimal.Decimal `json:"cbc_percent"`
	CbcValue                decimal.Decimal `json:"cbc_value"`
	TotalCannabinoidPercent decimal.Decimal `json:"total_cannabinoid_percent"`
	TotalCannabinoidValue   decimal.Decimal `json:"total_cannabinoid_value"`
}

type LabTestsPackage struct {
	LabTestID nulls.Int64 `json:"lab_test_id"`
	PackageID int64       `json:"package_id"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
}

type Order struct {
	ID                        int64           `json:"id"`
	CreatedAt                 time.Time       `json:"created_at"`
	UpdatedAt                 time.Time       `json:"updated_at"`
	ScheduledPackDateTime     time.Time       `json:"scheduled_pack_date_time"`
	ScheduledShipDateTime     time.Time       `json:"scheduled_ship_date_time"`
	ScheduledDeliveryDateTime time.Time       `json:"scheduled_delivery_date_time"`
	ActualPackDateTime        nulls.Time      `json:"actual_pack_date_time"`
	ActualShipDateTime        nulls.Time      `json:"actual_ship_date_time"`
	ActualDeliveryDateTime    nulls.Time      `json:"actual_delivery_date_time"`
	OrderTotal                decimal.Decimal `json:"order_total"`
	Notes                     string          `json:"notes"`
	Status                    string          `json:"status"`
	CustomerName              string          `json:"customer_name"`
}

type Package struct {
	ID                                int64           `json:"id"`
	CreatedAt                         time.Time       `json:"created_at"`
	UpdatedAt                         time.Time       `json:"updated_at"`
	TagID                             nulls.Int64     `json:"tag_id"`
	PackageType                       string          `json:"package_type"`
	IsActive                          bool            `json:"is_active"`
	Quantity                          decimal.Decimal `json:"quantity"`
	Notes                             string          `json:"notes"`
	PackagedDateTime                  time.Time       `json:"packaged_date_time"`
	HarvestDateTime                   nulls.Time      `json:"harvest_date_time"`
	LabTestingState                   string          `json:"lab_testing_state"`
	LabTestingStateDateTime           nulls.Time      `json:"lab_testing_state_date_time"`
	IsTradeSample                     bool            `json:"is_trade_sample"`
	IsTestingSample                   bool            `json:"is_testing_sample"`
	ProductRequiresRemediation        bool            `json:"product_requires_remediation"`
	ContainsRemediatedProduct         bool            `json:"contains_remediated_product"`
	RemediationDateTime               nulls.Time      `json:"remediation_date_time"`
	ReceivedDateTime                  nulls.Time      `json:"received_date_time"`
	ReceivedFromManifestNumber        nulls.String    `json:"received_from_manifest_number"`
	ReceivedFromFacilityLicenseNumber nulls.String    `json:"received_from_facility_license_number"`
	ReceivedFromFacilityName          nulls.String    `json:"received_from_facility_name"`
	IsOnHold                          bool            `json:"is_on_hold"`
	ArchivedDate                      nulls.Time      `json:"archived_date"`
	FinishedDate                      nulls.Time      `json:"finished_date"`
	ItemID                            nulls.Int64     `json:"item_id"`
	ProvisionalLabel                  nulls.String    `json:"provisional_label"`
	IsProvisional                     bool            `json:"is_provisional"`
	IsSold                            bool            `json:"is_sold"`
	PpuDefault                        decimal.Decimal `json:"ppu_default"`
	PpuOnOrder                        decimal.Decimal `json:"ppu_on_order"`
	TotalPackagePriceOnOrder          decimal.Decimal `json:"total_package_price_on_order"`
	PpuSoldPrice                      decimal.Decimal `json:"ppu_sold_price"`
	TotalSoldPrice                    decimal.Decimal `json:"total_sold_price"`
	PackagingSuppliesConsumed         bool            `json:"packaging_supplies_consumed"`
	IsLineItem                        bool            `json:"is_line_item"`
	OrderID                           nulls.Int64     `json:"order_id"`
	UomID                             int64           `json:"uom_id"`
	FacilityLocationID                int64           `json:"facility_location_id"`
}

type PackageAdjEntry struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	PackageID int64     `json:"package_id"`
	// can be negative or positive
	Amount decimal.Decimal `json:"amount"`
	UomID  int64           `json:"uom_id"`
}

type PackageAdjustment struct {
	ID            int64     `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	FromPackageID int64     `json:"from_package_id"`
	ToPackageID   int64     `json:"to_package_id"`
	// must be positive
	Amount decimal.Decimal `json:"amount"`
	UomID  int64           `json:"uom_id"`
}

type PackageTag struct {
	ID                int64       `json:"id"`
	CreatedAt         time.Time   `json:"created_at"`
	UpdatedAt         time.Time   `json:"updated_at"`
	TagNumber         string      `json:"tag_number"`
	IsAssigned        bool        `json:"is_assigned"`
	IsProvisional     bool        `json:"is_provisional"`
	IsActive          bool        `json:"is_active"`
	AssignedPackageID nulls.Int64 `json:"assigned_package_id"`
}

type ProductCategory struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type RetailerLocation struct {
	ID               int64           `json:"id"`
	CreatedAt        time.Time       `json:"created_at"`
	UpdatedAt        time.Time       `json:"updated_at"`
	Name             string          `json:"name"`
	Address          string          `json:"address"`
	City             string          `json:"city"`
	State            string          `json:"state"`
	Zip              string          `json:"zip"`
	Latitude         decimal.Decimal `json:"latitude"`
	Longitude        decimal.Decimal `json:"longitude"`
	Note             string          `json:"note"`
	Website          nulls.String    `json:"website"`
	SellsFlower      bool            `json:"sells_flower"`
	SellsPrerolls    bool            `json:"sells_prerolls"`
	SellsPressedHash bool            `json:"sells_pressed_hash"`
}

type Session struct {
	ID           uuid.UUID `json:"id"`
	Username     string    `json:"username"`
	RefreshToken string    `json:"refresh_token"`
	UserAgent    string    `json:"user_agent"`
	ClientIp     string    `json:"client_ip"`
	IsBlocked    bool      `json:"is_blocked"`
	ExpiresAt    time.Time `json:"expires_at"`
	CreatedAt    time.Time `json:"created_at"`
}

type SourcePackagesChildPackage struct {
	SourcePackageID int64     `json:"source_package_id"`
	ChildPackageID  int64     `json:"child_package_id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Strain struct {
	ID                      int64               `json:"id"`
	CreatedAt               time.Time           `json:"created_at"`
	UpdatedAt               time.Time           `json:"updated_at"`
	Name                    string              `json:"name"`
	Type                    string              `json:"type"`
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
	LightDep2022            string              `json:"light_dep_2022"`
	FallHarvest2022         string              `json:"fall_harvest_2022"`
	QuantityAvailable       decimal.Decimal     `json:"quantity_available"`
}

type Uom struct {
	ID           int64     `json:"id"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Name         string    `json:"name"`
	Abbreviation string    `json:"abbreviation"`
}

type User struct {
	ID                int64        `json:"id"`
	HashedPassword    string       `json:"hashed_password"`
	Username          string       `json:"username"`
	Email             string       `json:"email"`
	FirstName         string       `json:"first_name"`
	LastName          string       `json:"last_name"`
	Phone             nulls.String `json:"phone"`
	Role              string       `json:"role"`
	CreatedAt         time.Time    `json:"created_at"`
	PasswordChangedAt time.Time    `json:"password_changed_at"`
	UpdatedAt         time.Time    `json:"updated_at"`
}
