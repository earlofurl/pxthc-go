package seeds

import "github.com/rs/zerolog/log"

// productCategoriesSeed seeds the product_categories table with standard set of product categories
func (s Seed) productCategoriesSeed() {
	var err error

	_, err = s.tx.Exec(`INSERT INTO product_categories (name) VALUES ($1)`, "Buds")
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding product categories table at Buds")
	}

	_, err = s.tx.Exec(`INSERT INTO product_categories (name) VALUES ($1)`, "Buds (by Strain)")
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding product categories table at Buds (by Strain)")
	}

	_, err = s.tx.Exec(`INSERT INTO product_categories (name) VALUES ($1)`, "Concentrate")
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding product categories table at Concentrate")
	}

	_, err = s.tx.Exec(`INSERT INTO product_categories (name) VALUES ($1)`, "Concentrate (each)")
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding product categories table at Concentrate")
	}

	_, err = s.tx.Exec(`INSERT INTO product_categories (name) VALUES ($1)`, "Non-Infused (Plain) Pre-Roll")
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding product categories table at Non-Infused (Plain) Pre-Roll")
	}

	_, err = s.tx.Exec(`INSERT INTO product_categories (name) VALUES ($1)`, "Shake/Trim")
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding product categories table at Shake/Trim")
	}

	_, err = s.tx.Exec(`INSERT INTO product_categories (name) VALUES ($1)`, "Shake/Trim (by strain)")
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding product categories table at Shake/Trim (by strain)")
	}

	_, err = s.tx.Exec(`INSERT INTO product_categories (name) VALUES ($1)`, "Edible")
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding product categories table at Edible")
	}

	_, err = s.tx.Exec(`INSERT INTO product_categories (name) VALUES ($1)`, "Topical")
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding product categories table at Topical")
	}

	_, err = s.tx.Exec(`INSERT INTO product_categories (name) VALUES ($1)`, "Waste")
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding product categories table at Waste")
	}

	_, err = s.tx.Exec(`INSERT INTO product_categories (name) VALUES ($1)`, "Other")
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding product categories table at Other")
	}
}
