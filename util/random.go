package util

import (
	"fmt"
	"github.com/gobuffalo/nulls"
	"math/rand"
	"strings"
	"time"
)

// TODO: Make the random nulls.X values have a chance to return null

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

// RandomFloat generates a random decimal
func RandomFloat(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomNullsString generates a random nulls.String
func RandomNullsString(n int) nulls.String {
	return nulls.NewString(RandomString(n))
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomPhone generates a random phone number
func RandomPhone() nulls.String {
	return nulls.NewString(RandomString(10))
}

// RandomRole generates a random role
func RandomRole() string {
	roles := []string{"Admin", "User"}
	n := len(roles)
	return roles[rand.Intn(n)]
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

// RandomCategory generates a random category
func RandomCategory() string {
	return RandomString(6)
}

// RandomQuantity generates a random decimal quantity
func RandomQuantity() float64 {
	return RandomFloat(1, 20)
}

// RandomBool generates a random boolean value
func RandomBool() bool {
	return rand.Intn(2) == 0
}

// RandomNullsBool generates a random nulls.Bool
func RandomNullsBool() nulls.Bool {
	return nulls.NewBool(RandomBool())
}

// RandomPercent generates a random decimal percent
func RandomPercent() float64 {
	return RandomFloat(1, 100)
}

// RandomDecimalTimes100 generates a random decimal times 100
func RandomDecimalTimes100() float64 {
	return RandomFloat(0.01, 10.0) * 100
}

func RandomTagNumber() string {
	return RandomString(24)
}

// RandomValidEmail generates a random valid email address
func RandomValidEmail() string {
	return fmt.Sprintf("%s@%s.com", RandomString(6), RandomString(6))
}

// RandomTerpeneName returns a terpene name from a list of possible terpenes
func RandomTerpeneName() string {
	// slice of terpene names
	// TODO: Verify that the list is complete, and move the list to a constant for reuse
	terpenes := []string{
		"alpha-pinene",
		"beta-pinene",
		"limonene",
		"myrcene",
		"linalool",
		"caryophyllene",
		"terpinolene",
		"ocimene",
		"terpineol",
		"alpha-humulene",
		"beta-caryophyllene",
		"geraniol",
		"nerolidol",
		"alpha-terpinene",
		"alpha-phellandrene",
		"sabinene",
		"terpinene-4-ol",
		"alpha-bisabolol",
		"alpha-terpineol",
		"beta-phellandrene",
		"delta-3-carene",
		"delta-cadinene",
		"delta-cadinol",
		"delta-ocimene",
		"delta-terpinene",
		"gamma-terpinene",
		"geranyl-acetate",
		"geranyl-acetone",
		"geranyl-butyrate",
		"geranyl-geraniol",
		"geranyl-isobutyrate",
		"geranyl-isovalerate",
		"geranyl-terpinol",
	}
	n := len(terpenes)
	return terpenes[rand.Intn(n)]
}

// RandomStrainType returns a strain type from a list of possible strain types
func RandomStrainType() string {
	// slice of strain types
	strainTypes := []string{"Indica", "Sativa", "Hybrid", "Indica Dominant", "Sativa Dominant"}
	n := len(strainTypes)
	return strainTypes[rand.Intn(n)]
}

// RandomDateTime returns a random date time within the last 30 days
func RandomDateTime() time.Time {
	return time.Now().AddDate(0, 0, -rand.Intn(30)).Round(time.Microsecond)
}
