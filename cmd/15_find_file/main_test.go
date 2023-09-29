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
		{
			In:  "1.avi\n12\nemoh\n vonavi\n  a.doc\n  b.doc \n vortep\n  .bashrc\n vorodis\n  onrop\n   1.avi\n   2.avi \nrav\n bil\n",
			Out: "/emoh/vorodis/onrop/1.avi\n",
		},
		{
			In:  "1.avi\n2\nemoh\n1.avi\n",
			Out: "/1.avi\n",
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
