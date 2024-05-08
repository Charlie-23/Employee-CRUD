package util

import (
	"fmt"
	"math/rand"
	"strings"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// func init() {
// 	rand.Seed(time.Now().UnixNano())
// }

// RandomInt generates a random integer between min and max
func RandomFloat() float64 {
	// Define the range
	min := 0.0
	max := 10000000.0

	// Generate a random float within the range
	randomFloat := min + rand.Float64()*(max-min)
	return randomFloat
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

// RandomOwner generates a random owner name
func RandomName() string {
	return RandomString(6)
}

// RandomMoney generates a random amount of money
func RandomSalary() float64 {
	return RandomFloat()
}

// RandomCurrency generates a random currency code
func RandomPosition() string {
	positions := []string{"SDE", "SDE II", "SDE III", "SWE Manager", "PO", "GPO"}
	n := len(positions)
	return positions[rand.Intn(n)]
}

// RandomEmail generates a random email
func RandomEmail() string {
	return fmt.Sprintf("%s@email.com", RandomString(6))
}

func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1)
}

