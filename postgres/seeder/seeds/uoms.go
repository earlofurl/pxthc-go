package seeds

import "github.com/rs/zerolog/log"

// uomsSeed seeds the uoms table with standard set of units of measure
func (s Seed) uomsSeed() {
	var err error

	_, err = s.tx.Exec(`INSERT INTO uoms (name, abbreviation) VALUES ($1, $2)`, "Each", "ea")
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding uoms table at Each")
	}

	_, err = s.tx.Exec(`INSERT INTO uoms (name, abbreviation) VALUES ($1, $2)`, "Pounds", "lb")
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding uoms table at Pounds")
	}

	_, err = s.tx.Exec(`INSERT INTO uoms (name, abbreviation) VALUES ($1, $2)`, "Grams", "g")
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding uoms table at Grams")
	}

	_, err = s.tx.Exec(`INSERT INTO uoms (name, abbreviation) VALUES ($1, $2)`, "Kilograms", "kg")
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding uoms table at Kilograms")
	}

	_, err = s.tx.Exec(`INSERT INTO uoms (name, abbreviation) VALUES ($1, $2)`, "Milligrams", "mg")
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding uoms table at Milligrams")
	}

	_, err = s.tx.Exec(`INSERT INTO uoms (name, abbreviation) VALUES ($1, $2)`, "Ounces", "oz")
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding uoms table at Ounces")
	}
}
