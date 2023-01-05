package seeds

import (
	"fmt"
	"github.com/earlofurl/pxthc/util"
	"github.com/rs/zerolog/log"
)

// packagesSeed seeds the packages table with standard set of packages
func (s Seed) packagesSeed() {
	var err error

	for i := 0; i < 100; i++ {
		i := i
		_, err = s.tx.Exec(`INSERT INTO packages (
                      tag_id, 
                      package_type,
                      is_active,
                      quantity,
                      notes,
                      packaged_date_time,
                      harvest_date_time,
                      lab_testing_state,
                      lab_testing_state_date_time,
                      is_trade_sample,
                      is_testing_sample,
                      product_requires_remediation,
                      contains_remediated_product,
                      remediation_date_time,
                      received_date_time,
                      received_from_manifest_number,
                      received_from_facility_license_number,
                      received_from_facility_name,
                      is_on_hold,
                      archived_date,
                      finished_date,
                      item_id,
                      provisional_label,
                      is_provisional,
                      is_sold,
                      ppu_default,
                      ppu_on_order,
                      total_package_price_on_order,
                      ppu_sold_price,
                      total_sold_price,
                      packaging_supplies_consumed,
                      is_line_item,
                      order_id,
                      uom_id,
                      facility_location_id
                      ) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35)`,
			i+1,
			fmt.Sprintf("Package Type %d", i+1),
			true,
			util.RandomFloat(0, 100),
			fmt.Sprintf("Notes %d", i+1),
			util.RandomDateTime(),
			util.RandomDateTime(),
			fmt.Sprintf("Lab Testing State %d", i+1),
			util.RandomDateTime(),
			false,
			false,
			false,
			false,
			nil,
			nil,
			nil,
			nil,
			nil,
			false,
			nil,
			nil,
			util.RandomInt(1, 100),
			fmt.Sprintf("Provisional Label %d", i+1),
			false,
			false,
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			false,
			false,
			nil,
			util.RandomInt(1, 6),
			util.RandomInt(1, 10))
		if err != nil {
			log.Fatal().Err(err).Msg("Error seeding packages table")
		}
	}
}
