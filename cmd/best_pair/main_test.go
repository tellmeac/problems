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
		{
			In: `6
10 3 5 3 11 9
`,
			Out: `2 5
`,
		},
		//		{
		//			In: `4
		//4500 1024 3043 4980
		//`,
		//			Out: `2 4
		//`,
		//		},
		{
			In: `8
8 12 24 2 3 8 10 15
`,
			Out: `4 8
`,
		},
		{
			In: `1
10
`,
			Out: `0 0
`,
		},
		{
			In: `5
1034 1 1001 1999 5000
`,
			Out: `2 5
`,
		},
		{
			In: `11
10 15 7 8 2 7 4 3 7 10 12
`,
			Out: `5 11
`,
		},
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
