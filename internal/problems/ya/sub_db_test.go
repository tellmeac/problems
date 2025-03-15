package ya

import (
	"io"
	"strings"
	"testing"
	"time"

	"github.com/tellmeac/problems/pkg/std"
)

func TestSubscribeDatabase(t *testing.T) {
	std.New(func(r io.Reader, w io.Writer) error {
		SubscribeDatabase(r, w)
		return nil
	}).WithTimeout(time.Second).
		WithData(strings.NewReader(`2 5
2 0 price stock_count
1 0 partner_content
{"trace_id":"1","offer":{"id":"1","price":9990}}
{"trace_id":"2","offer":{"id":"1","stock_count":100}}
{"trace_id":"3","offer":{"id":"2","partner_content":{"title":"Backpack"}}}
{"trace_id":"4","offer":{"id":"1","stock_count":100}}
{"trace_id":"5","offer":{"id":"2","partner_content":{"title":"Backpack"}}}`)).
		WithValidator(std.ValidateStrict(`{"trace_id":"1","offer":{"id":"1","price":9990}}
{"trace_id":"2","offer":{"id":"1","price":9990,"stock_count":100}}
{"trace_id":"3","offer":{"id":"2","partner_content":{"title":"Backpack"}}}
{"trace_id":"4","offer":{"id":"1","price":9990,"stock_count":100}}
{"trace_id":"5","offer":{"id":"2","partner_content":{"title":"Backpack"}}}
`)).Run(t)
}
