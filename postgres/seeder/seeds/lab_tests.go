package seeds

import (
	"fmt"
	"github.com/earlofurl/pxthc/util"
	"github.com/rs/zerolog/log"
)

// labTestsSeed seeds the lab_tests table with standard set of lab tests
func (s Seed) labTestsSeed() {
	var err error

	for i := 0; i < 100; i++ {
		i := i
		_, err = s.tx.Exec(`INSERT INTO lab_tests (
                       test_name, 
                       batch_code, 
                       test_id_code, 
                       lab_facility_name, 
                       test_performed_date_time, 
                       test_completed, 
                       overall_passed, 
                       test_type_name, 
                       test_passed, 
                       test_comment, 
                       thc_total_percent, 
                       thc_total_value, 
                       cbd_percent, 
                       cbd_value, 
                       terpene_total_percent, 
                       terpene_total_value, 
                       thc_a_percent, 
                       thc_a_value, 
                       delta9_thc_percent, 
                       delta9_thc_value, 
                       delta8_thc_percent, 
                       delta8_thc_value, 
                       thc_v_percent, 
                       thc_v_value, 
                       cbd_a_percent, 
                       cbd_a_value, 
                       cbn_percent, 
                       cbn_value, 
                       cbg_a_percent, 
                       cbg_a_value, 
                       cbg_percent, 
                       cbg_value, 
                       cbc_percent, 
                       cbc_value, 
                       total_cannabinoid_percent, 
                       total_cannabinoid_value) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35, $36)`,
			fmt.Sprintf("Lab Test %d", i+1),
			fmt.Sprintf("%d-%s-%s", util.RandomInt(20, 23), util.RandomString(2), util.RandomString(1)),
			fmt.Sprintf("Test Code %d", i+1),
			fmt.Sprintf("Facility %d", i+1),
			util.RandomDateTime(),
			true,
			true,
			fmt.Sprintf("Test Type %d", i+1),
			true,
			fmt.Sprintf("Test Comment %d", i+1),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100),
			util.RandomFloat(0, 100))
		if err != nil {
			log.Fatal().Err(err).Msg("Error seeding lab tests table")
		}
	}
}
