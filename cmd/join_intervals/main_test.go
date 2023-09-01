package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTask(t *testing.T) {
	for i, tc := range []struct {
		In  []int
		Out string
	}{
		{In: []int{1, 2, 3, 5, 6, 7}, Out: "1-3,5-7"},
		{In: []int{1, 2, 3, 5, 6, 8}, Out: "1-3,5-6,8"},
		{In: []int{1}, Out: "1"},
		{In: []int{-2, -3, 5}, Out: "-3--2,5"},
	} {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			res := compress(tc.In)

			assert.Equal(t, tc.Out, res)
		})
	}
}
