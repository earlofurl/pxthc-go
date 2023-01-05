package seeds

import "github.com/rs/zerolog/log"

// usersSeed seeds the users table with standard set of users
func (s Seed) usersSeed() {
	var err error

	// password is "secret"
	_, err = s.tx.Exec(`INSERT INTO users (hashed_password, username, email, first_name, last_name, phone, role) 
			VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		"$2a$10$pjIDTs0JPoxVpp22QASjvusw/0aUA5l0QgfcFSqL3X4p6yZow7zMa",
		"test-user-admin-01",
		"test@email.com",
		"Test",
		"User",
		"1234567890",
		"admin")
	if err != nil {
		log.Fatal().Err(err).Msg("Error seeding users table at test-user-admin-01")
	}
}
