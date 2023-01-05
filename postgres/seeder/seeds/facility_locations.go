package seeds

import (
	"fmt"
	"github.com/earlofurl/pxthc/util"
	"github.com/rs/zerolog/log"
)

// facilityLocationsSeed seeds the facility_locations table with standard set of facility locations
func (s Seed) facilityLocationsSeed() {
	var err error

	for i := 0; i < 100; i++ {
		i := i
		_, err = s.tx.Exec(`INSERT INTO facility_locations (name, facility_id) 
			VALUES ($1, $2)`,
			fmt.Sprintf("Facility Location %d", i+1),
			util.RandomInt(1, 100))
		if err != nil {
			log.Fatal().Err(err).Msg("Error seeding facility locations table")
		}
	}

}
