package seeds

import (
	"fmt"
	"github.com/earlofurl/pxthc/util"
	"github.com/rs/zerolog/log"
	"strconv"
)

// facilitiesSeed seeds the facilities table with standard set of facilities
func (s Seed) facilitiesSeed() {
	var err error

	for i := 0; i < 100; i++ {
		i := i
		_, err = s.tx.Exec(`INSERT INTO facilities (name, license_number) 
		VALUES ($1, $2)`,
			fmt.Sprintf("Facility %d", i+1),
			strconv.FormatInt(util.RandomInt(100000, 999999), 10))
		if err != nil {
			log.Fatal().Err(err).Msg("Error seeding facilities table")
		}
	}

}
