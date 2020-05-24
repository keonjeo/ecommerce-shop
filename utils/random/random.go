package random

import (
	"math/rand"
)

// IntFromRange returns random int within range from min to max
func IntFromRange(min, max int) int {
	return rand.Intn((max-min)+1) + min
}
