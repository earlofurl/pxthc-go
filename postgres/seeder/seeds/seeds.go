package seeds

import (
	"github.com/jmoiron/sqlx"

	// postgres driver.
	_ "github.com/lib/pq"
)

type Seed struct {
	tx *sqlx.Tx
}

// PopulateDatabase calls the seed functions in a specific order
func (s Seed) PopulateDatabase() {
	s.uomsSeed()
	s.facilitiesSeed()
	s.facilityLocationsSeed()
	s.usersSeed()
	s.strainsSeed()
	s.productCategoriesSeed()
	s.itemTypesSeed()
	s.packageTagsSeed()
	s.itemsSeed()
	s.labTestsSeed()
	s.packagesSeed()
	s.labTestsPackagesSeed()
}

// NewSeed returns a Seed with a pool of connections to the database
func NewSeed(tx *sqlx.Tx) Seed {
	return Seed{
		tx: tx,
	}
}
