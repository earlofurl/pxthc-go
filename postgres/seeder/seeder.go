package main

import (
	"github.com/earlofurl/pxthc/postgres/seeder/seeds"
	"github.com/jmoiron/sqlx"

	"github.com/danvergara/seeder"
	"github.com/rs/zerolog/log"

	// postgres driver
	_ "github.com/lib/pq"
)

func main() {
	db, err := sqlx.Open("postgres", "postgres://root:secret@localhost:5432/pixel_thc_dev?sslmode=disable")
	if err != nil {
		log.Fatal().Err(err).Msg("Error connecting to database")
	}

	tx, err := db.Beginx()
	if err != nil {
		log.Fatal().Err(err).Msg("Error starting transaction")
	}

	// Create a new seeder
	s := seeds.NewSeed(tx)

	if err := seeder.Execute(s); err != nil {
		log.Fatal().Err(err).Msg("Error seeding database. Rolling back...")
		err := tx.Rollback()
		if err != nil {
			return
		}
	}
	err = tx.Commit()
	if err != nil {
		return
	}
}
