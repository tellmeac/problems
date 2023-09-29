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
		{In: "5\n1\n0\n1\n0\n1\n", Out: "1\n"},
		{In: "5\n1\n1\n1\n0\n1\n", Out: "3\n"},
		{In: "5\n1\n0\n1\n1\n1\n", Out: "3\n"},
		{In: "5\n0\n0\n1\n1\n0\n", Out: "2\n"},
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
