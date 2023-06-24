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
		{In: "5\n1 2 3 4 5\n", Out: "3\n"},
		{In: "10\n1 1 2 3 5 5 2 2 1 5\n", Out: "4\n"},
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
