package seeds

import (
	"fmt"
	"github.com/earlofurl/pxthc/util"
	"github.com/rs/zerolog/log"
)

// strainsSeed seeds the strains table with standard set of strains
func (s Seed) strainsSeed() {
	var err error

	for i := 0; i < 100; i++ {
		i := i
		_, err = s.tx.Exec(`INSERT INTO strains (name, type, yield_average, terp_average_total, terp_1, terp_1_value, terp_2, terp_2_value, terp_3, terp_3_value, terp_4, terp_4_value, terp_5, terp_5_value, thc_average, total_cannabinoid_average, light_dep_2022, fall_harvest_2022, quantity_available) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)`,
			fmt.Sprintf("Strain %d", i+1),
			util.RandomStrainType(),
			util.RandomFloat(0.001, 20),
			util.RandomFloat(0.001, 20),
			util.RandomTerpeneName(),
			util.RandomFloat(0.001, 20),
			util.RandomTerpeneName(),
			util.RandomFloat(0.001, 20),
			util.RandomTerpeneName(),
			util.RandomFloat(0.001, 20),
			util.RandomTerpeneName(),
			util.RandomFloat(0.001, 20),
			util.RandomTerpeneName(),
			util.RandomFloat(0.001, 20),
			util.RandomFloat(0.001, 20),
			util.RandomFloat(0.001, 20),
			util.RandomBool(),
			util.RandomBool(),
			util.RandomFloat(0.001, 20))
		if err != nil {
			log.Fatal().Err(err).Msg("Error seeding strains table")
		}
	}
}
