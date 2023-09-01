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
		{In: "5\n00:00:00\n00:01:11\n02:15:59\n23:59:58\n23:59:59", Out: "1"},
		{In: "3\n12:00:00\n23:59:59\n00:00:00\n", Out: "2"},
		{In: "4\n00:00:00\n00:00:00\n00:00:00\n00:00:00\n", Out: "4"},
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
