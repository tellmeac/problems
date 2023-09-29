package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestTask(t *testing.T) {
	for i, tc := range []struct {
		In, Out string
	}{
		{In: "22 18 8\n6 11\n10 15\n3 18\n1 19\n10 17\n1 10\n6 16\n20 21\n1 1\n12 21\n5 9\n1 10\n5 10\n6 11\n5 6\n7 11\n1 19\n13 15\n5\n22\n19\n3\n8\n16\n16\n21\n", Out: "8\n0\n3\n5\n11\n6\n6\n2\n"},
	} {
		t.Run(fmt.Sprintf("case %d", i), func(t *testing.T) {
			reader = bufio.NewReader(strings.NewReader(tc.In))

			buffer := bytes.NewBufferString("")
			writer = bufio.NewWriter(buffer)

			main()

			_ = writer.Flush()

			assert.Equal(t, tc.Out, buffer.String())
			t.Log(buffer.String())
		})
	}
}
