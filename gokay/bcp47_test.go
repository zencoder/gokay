package gokay

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIsBCP47_Nil(t *testing.T) {
	require.NoError(t, IsBCP47(nil))
}

func TestIsBCP47_English(t *testing.T) {
	str := "English"
	require.EqualError(t, IsBCP47(&str), "language: tag is not well-formed")
}

func TestIsBCP47_en(t *testing.T) {
	str := "en"
	require.NoError(t, IsBCP47(&str))
}

func TestIsBCP47_en_AB(t *testing.T) {
	str := "en-AB"
	require.NoError(t, IsBCP47(&str))
}

func TestIsBCP47OrEmpty_Nil(t *testing.T) {
	require.NoError(t, IsBCP47OrEmpty(nil))
}

func TestIsBCP47OrEmpty_English(t *testing.T) {
	str := "English"
	require.EqualError(t, IsBCP47OrEmpty(&str), "language: tag is not well-formed")
}

func TestIsBCP47OrEmpty_en(t *testing.T) {
	str := "en"
	require.NoError(t, IsBCP47OrEmpty(&str))
}

func TestIsBCP47OrEmpty_en_AB(t *testing.T) {
	str := "en-AB"
	require.NoError(t, IsBCP47OrEmpty(&str))
}

func TestIsBCP47OrEmpty_empty(t *testing.T) {
	str := ""
	require.NoError(t, IsBCP47OrEmpty(&str))
}
