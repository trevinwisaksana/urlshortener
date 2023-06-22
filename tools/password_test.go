package tools

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func createPasswordHash(t *testing.T) string {
	password := RandomAlphanumericString(6)

	hashedPassword, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword)

	err = ComparePasswordHash(password, hashedPassword)
	require.NoError(t, err)

	return hashedPassword
}

func TestHashPassword(t *testing.T) {
	createPasswordHash(t)
}

func TestIncorrectPasswordHash(t *testing.T) {
	passwordHash := createPasswordHash(t)

	wrongPassword := RandomAlphanumericString(6)
	err := ComparePasswordHash(wrongPassword, passwordHash)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())
}
