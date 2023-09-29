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
		{In: "5\n2\n4\n8\n8\n8\n", Out: "2\n4\n8\n"},
		{In: "5\n2\n4\n8\n", Out: "2\n4\n8\n"},
		{In: "5\n2\n2\n8\n", Out: "2\n8\n"},
		{In: "5\n2\n2\n", Out: "2\n"},
		{In: "0\n", Out: ""},
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
