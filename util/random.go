package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

// init() function will be called automatically when the package is first used.
// if use rand.Seed(1) will use constant value
func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
// rand.Int63(max-min+1) -> 0~max-min
// min ~max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63(max-min+1)
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
