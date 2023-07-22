package main

import (
	"bufio"
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTask(t *testing.T) {
	for i, tc := range []struct {
		In, Out string
	}{
		{In: "3\n17 16 18", Out: "2\n"},
		{In: "5\n120 25 30 100 105 ", Out: "4\n"},
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
