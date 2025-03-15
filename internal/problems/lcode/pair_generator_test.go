package lcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPairGenerator(t *testing.T) {
	want := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	got := PairGenerator(3)

	require.ElementsMatch(t, want, got)
}
