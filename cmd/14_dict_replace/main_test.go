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
		{In: "a b\nabdafb basrt casds dsasa a", Out: "a b casds dsasa a\n"},
		{In: "aa bc aaa\na aa aaa bcd abcd\n", Out: "a aa aa bc abcd\n"},
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
