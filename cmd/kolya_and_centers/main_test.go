package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestKolya(t *testing.T) {
	for i, tc := range []struct {
		Input    string
		Expected string
	}{
		{
			Input: `3 3 12
DISABLE 1 2
DISABLE 2 1
DISABLE 3 3
GETMAX
RESET 1
RESET 2
DISABLE 1 2
DISABLE 1 3
DISABLE 2 2
GETMAX
RESET 3
GETMIN
`,
			Expected: `1
2
1
`,
		},
		{
			Input: `2 3 9
DISABLE 1 1
DISABLE 2 2
RESET 2
DISABLE 2 1
DISABLE 2 3
RESET 1
GETMAX
DISABLE 2 1
GETMIN`,
			Expected: `1
2
`,
		},
	} {
		t.Run(fmt.Sprintf("Test Case %d", i), func(t *testing.T) {
			reader = bufio.NewReader(strings.NewReader(tc.Input))

			answerWriter := bytes.NewBufferString("")
			writer = bufio.NewWriter(answerWriter)

			main()

			_ = writer.Flush()
			assert.Equal(t, tc.Expected, answerWriter.String())

			t.Log(answerWriter.String())
		})
	}
}

func TestDataCenter_Less(t *testing.T) {
	a := &DataCenter{
		ID:       1,
		Restarts: 0,
		Active:   3,
	}
	b := &DataCenter{
		ID:       3,
		Restarts: 0,
		Active:   3,
	}
	c := &DataCenter{
		ID:       2,
		Restarts: 1,
		Active:   3,
	}

	assert.Truef(t, a.Greater(b), "By less identifier")
	assert.Falsef(t, b.Greater(a), "By greater identifier")
	assert.Truef(t, c.Greater(a), "By value")
	assert.Truef(t, c.Greater(b), "By value")
	assert.Falsef(t, a.Greater(c), "By value")
	assert.Falsef(t, b.Greater(c), "By value")
}
