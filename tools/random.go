package tools

import (
	"math/rand"
	"strings"
	"time"
)

const alphanumeric = "abcdefghijklmnopqrstuvwxyz1234567890-_"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func hasUnderscoreSuffix(r string) bool {
	if strings.HasSuffix(r, "-") || strings.HasSuffix(r, "_") {
		return true
	}

	return false
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphanumeric)

	for i := 0; i < n; i++ {
		c := alphanumeric[rand.Intn(k)]
		sb.WriteByte(c)
	}

	result := sb.String()

	if hasUnderscoreSuffix(result) {
		return RandomString(n)
	}

	return sb.String()
}

func RandomUsername() string {
	return RandomString(6)
}

func RandomName() string {
	return RandomString(6)
}

// RandomEmail generates a random email
func RandomEmail() string {
	return RandomString(5) + "@email.com"
}
