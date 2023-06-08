package tools

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandomString(t *testing.T) {
	randomID := RandomString(5)

	require.Equal(t, len(randomID), 5)
	require.NotNil(t, randomID)
}

func TestHasUnderscoreSuffix(t *testing.T) {
	sample := "abcdefghijklmnopqrstuvwxyz1234567890-_"
	result := hasUnderscoreSuffix(sample)

	require.True(t, result)
}

func TestHasNoUnderscoreSuffix(t *testing.T) {
	sample := "abcdefghijklmnopqrstuvwxyz1234567890"
	result := hasUnderscoreSuffix(sample)

	require.False(t, result)
}
