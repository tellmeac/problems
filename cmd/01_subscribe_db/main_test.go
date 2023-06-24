package main

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestSubscribeDB(t *testing.T) {
	for i, tc := range []struct {
		Input    string
		Expected string
	}{
		{
			Input: `2 5
2 0 price stock_count
1 0 partner_content
{"trace_id":"1","offer":{"id":"1","price":9990}}
{"trace_id":"2","offer":{"id":"1","stock_count":100}}
{"trace_id":"3","offer":{"id":"2","partner_content":{"title":"Backpack"}}}
{"trace_id":"4","offer":{"id":"1","stock_count":100}}
{"trace_id":"5","offer":{"id":"2","partner_content":{"title":"Backpack"}}}`,
			Expected: `{"trace_id":"1","offer":{"id":"1","price":9990}}
{"trace_id":"2","offer":{"id":"1","price":9990,"stock_count":100}}
{"trace_id":"3","offer":{"id":"2","partner_content":{"title":"Backpack"}}}
{"trace_id":"4","offer":{"id":"1","price":9990,"stock_count":100}}
{"trace_id":"5","offer":{"id":"2","partner_content":{"title":"Backpack"}}}
`,
		},
		{
			Input: `1 2
1 0 title
{"trace_id":"1","offer":{"id":"1","partner_content":{"title":"Backpack"}}}
{"trace_id":"2","offer":{"id":"1","partner_content":{"description":"Backpack description"}}}`,
			Expected: `{"trace_id":"1","offer":{"id":"1","partner_content":{"title":"Backpack"}}}
`,
		},
		{
			Input: `3 3
1 0 price
1 1 price description
2 0 stock_count price
{"trace_id":"1","offer":{"id":"1","partner_content":{"description":"Backpack description"}}}
{"trace_id":"2","offer":{"id":"1","price":9990}}
{"trace_id":"3","offer":{"id":"1","stock_count":100}}`,
			Expected: `{"trace_id":"2","offer":{"id":"1","price":9990}}
{"trace_id":"2","offer":{"id":"1","partner_content":{"description":"Backpack description"},"price":9990}}
{"trace_id":"2","offer":{"id":"1","price":9990}}
{"trace_id":"3","offer":{"id":"1","price":9990,"stock_count":100}}
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
