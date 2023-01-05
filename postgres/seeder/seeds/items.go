package seeds

import (
	"fmt"
	"github.com/rs/zerolog/log"
)

// itemsSeed seeds the items table with standard set of items
func (s Seed) itemsSeed() {
	var err error

	for i := 0; i < 100; i++ { // cycle through 100 strains
		i := i
		for j := 0; j < 10; j++ { // cycle through 10 items per strain
			j := j
			_, err = s.tx.Exec(`INSERT INTO items (description, is_used, item_type_id, strain_id) 
			VALUES ($1, $2, $3, $4)`,
				fmt.Sprintf("Item S-%d T-%d", i+1, j+1),
				false,
				j+1,
				i+1)
			if err != nil {
				log.Fatal().Err(err).Msg("Error seeding items table")
			}
		}
	}
}
