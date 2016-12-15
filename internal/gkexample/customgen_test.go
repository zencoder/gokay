package gkexample

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCustomGenerator(t *testing.T) {
	require.NotNil(t, NewCustomGenerator())
}
