package tools

import (
	"math/rand"
	"strings"
	"time"
)

const alphanumeric = "abcdefghijklmnopqrstuvwxyz1234567890-_"
const alphabets = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func hasUnderscoreSuffix(r string) bool {
	if strings.HasSuffix(r, "-") || strings.HasSuffix(r, "_") {
		return true
	}

	return false
}

// RandomAlphanumericString generates a random string of length n
func RandomAlphanumericString(n int) string {
	var sb strings.Builder
	k := len(alphanumeric)

	for i := 0; i < n; i++ {
		c := alphanumeric[rand.Intn(k)]
		sb.WriteByte(c)
	}

	result := sb.String()

	if hasUnderscoreSuffix(result) {
		return RandomAlphanumericString(n)
	}

	return sb.String()
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabets)

	for i := 0; i < n; i++ {
		c := alphabets[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomUsername() string {
	return RandomString(6)
}

func RandomName() string {
	return RandomAlphanumericString(6)
}

// RandomEmail generates a random email
func RandomEmail() string {
	return RandomAlphanumericString(5) + "@email.com"
}
