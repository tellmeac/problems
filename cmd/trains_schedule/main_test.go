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
		{In: "3\n00:00 23:59 00:00", Out: "0\n"},
		{In: "2\n23:59 00:00", Out: "1\n"},
		{In: "4\n12:23 23:49 04:20 16:48", Out: "265\n"},
		{In: "1\n12:23", Out: "1440\n"},
		{In: "1\n01:23", Out: "1440\n"},
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
