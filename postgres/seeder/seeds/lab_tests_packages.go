package seeds

import (
	"github.com/earlofurl/pxthc/util"

	"github.com/rs/zerolog/log"
)

// labTestsPackagesSeed seeds the lab_tests_packages junction table with standard set of lab tests and packages
func (s Seed) labTestsPackagesSeed() {
	var err error

	for i := 0; i < 100; i++ {
		i := i
		_, err = s.tx.Exec(`INSERT INTO lab_tests_packages (lab_test_id, package_id) 
			VALUES ($1, $2)`,
			util.RandomInt(1, 100),
			i+1)
		if err != nil {
			log.Fatal().Err(err).Msg("Error seeding lab tests packages table")
		}
	}
}
