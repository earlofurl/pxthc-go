package seeds

import (
	"fmt"
	"github.com/rs/zerolog/log"
)

// TODO: Speed up this seed by using a single query to insert multiple rows

// packageTagsSeed seeds the package_tags table with standard set of package tags
func (s Seed) packageTagsSeed() {
	var err error
	var packageTagPrefix = "1A4010300016BAD0000"

	for i := 10000; i < 20000; i++ {
		i := i
		_, err = s.tx.Exec(`INSERT INTO package_tags (tag_number, is_assigned, is_provisional, is_active, assigned_package_id) 
				VALUES ($1, $2, $3, $4, $5)`,
			fmt.Sprintf("%s%04d", packageTagPrefix, i),
			false,
			false,
			false,
			nil)
		if err != nil {
			log.Fatal().Err(err).Msg("Error seeding package tags table")
		}
	}
}
