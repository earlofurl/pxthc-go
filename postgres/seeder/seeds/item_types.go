package seeds

import (
	"github.com/rs/zerolog/log"
)

// itemTypesSeed seeds the item_types table with standard set of item types
func (s Seed) itemTypesSeed() {
	var err error
	var uomPoundsID int64
	var uomGramsID int64
	var uomEachID int64
	var prodCatBudsID int64
	var prodCatPrerollsID int64
	var prodCatConcentrateID int64
	var prodCatConcentrateEachID int64

	// Get IDs for previously seeded UoMs
	err = s.tx.Get(&uomPoundsID, `SELECT id FROM uoms WHERE name = 'Pounds'`)
	if err != nil {
		log.Fatal().Err(err).Msg("Error getting ID from uom table for item types table at Pounds")
	}

	err = s.tx.Get(&uomGramsID, `SELECT id FROM uoms WHERE name = 'Grams'`)
	if err != nil {
		log.Fatal().Err(err).Msg("Error getting ID from uom table for item types table at Grams")
	}

	err = s.tx.Get(&uomEachID, `SELECT id FROM uoms WHERE name = 'Each'`)
	if err != nil {
		log.Fatal().Err(err).Msg("Error getting ID from uom table for item types table at Each")
	}

	// Get IDs for previously seeded Product Categories
	err = s.tx.Get(&prodCatBudsID, `SELECT id FROM product_categories WHERE name = 'Buds'`)
	if err != nil {
		log.Fatal().Err(err).Msg("Error getting ID from product category table for item types table at Buds")
	}

	err = s.tx.Get(&prodCatPrerollsID, `SELECT id FROM product_categories WHERE name = 'Non-Infused (Plain) Pre-Roll'`)
	if err != nil {
		log.Fatal().Err(err).Msg("Error getting ID from product category table for item types table at Non-Infused (Plain) Pre-Roll")
	}

	err = s.tx.Get(&prodCatConcentrateID, `SELECT id FROM product_categories WHERE name = 'Concentrate'`)
	if err != nil {
		log.Fatal().Err(err).Msg("Error getting ID from product category table for item types table at Concentrate")
	}

	err = s.tx.Get(&prodCatConcentrateEachID, `SELECT id FROM product_categories WHERE name = 'Concentrate (each)'`)
	if err != nil {
		log.Fatal().Err(err).Msg("Error getting ID from product category table for item types table at Concentrate (each)")
	}

	// Insert values
	_, err = s.tx.Exec(`INSERT INTO item_types (product_form, product_modifier, uom_default, product_category_id) 
			VALUES ($1, $2, $3, $4)`,
		"Flower", "A Bud", uomGramsID, prodCatBudsID)
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding item types table at Flower A Bud")
	}

	_, err = s.tx.Exec(`INSERT INTO item_types (product_form, product_modifier, uom_default, product_category_id) 
			VALUES ($1, $2, $3, $4)`,
		"Flower", "B Bud", uomGramsID, prodCatBudsID)
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding item types table at Flower B Bud")
	}

	_, err = s.tx.Exec(`INSERT INTO item_types (product_form, product_modifier, uom_default, product_category_id) 
			VALUES ($1, $2, $3, $4)`,
		"Flower", "Unsorted", uomPoundsID, prodCatBudsID)
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding item types table at Flower Unsorted")
	}

	_, err = s.tx.Exec(`INSERT INTO item_types (product_form, product_modifier, uom_default, product_category_id) 
			VALUES ($1, $2, $3, $4)`,
		"Flower", "Mini", uomGramsID, prodCatBudsID)
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding item types table at Flower Mini")
	}

	_, err = s.tx.Exec(`INSERT INTO item_types (product_form, product_modifier, uom_default, product_category_id) 
			VALUES ($1, $2, $3, $4)`,
		"Flower", "Ground PRM", uomGramsID, prodCatBudsID)
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding item types table at Flower Ground PRM")
	}

	_, err = s.tx.Exec(`INSERT INTO item_types (product_form, product_modifier, uom_default, product_category_id) 
			VALUES ($1, $2, $3, $4)`,
		"Preroll", "Single", uomGramsID, prodCatBudsID)
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding item types table at Preroll Single")
	}

	_, err = s.tx.Exec(`INSERT INTO item_types (product_form, product_modifier, uom_default, product_category_id) 
			VALUES ($1, $2, $3, $4)`,
		"Preroll", "2-pack", uomEachID, prodCatPrerollsID)
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding item types table at Preroll 2-pack")
	}

	_, err = s.tx.Exec(`INSERT INTO item_types (product_form, product_modifier, uom_default, product_category_id) 
			VALUES ($1, $2, $3, $4)`,
		"Preroll", "10-pack", uomEachID, prodCatPrerollsID)
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding item types table at Preroll 10-pack")
	}

	_, err = s.tx.Exec(`INSERT INTO item_types (product_form, product_modifier, uom_default, product_category_id) 
			VALUES ($1, $2, $3, $4)`,
		"Hash", "Bulk", uomGramsID, prodCatConcentrateID)
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding item types table at Hash Bulk")
	}

	_, err = s.tx.Exec(`INSERT INTO item_types (product_form, product_modifier, uom_default, product_category_id) 
			VALUES ($1, $2, $3, $4)`,
		"Hash", "Packaged", uomEachID, prodCatConcentrateEachID)
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding item types table at Hash Packaged")
	}
}
