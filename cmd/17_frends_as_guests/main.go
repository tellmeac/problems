package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

type event struct {
	friend, start, end int
}

func (e event) Days() int {
	return e.end - e.start
}

func main() {
	defer func() {
		_ = writer.Flush()
	}()

	var n int
	read("%d\n", &n)

	events := make([]event, n)
	for idx := 0; idx < n; idx++ {
		events[idx].friend = idx
		read("%d %d\n", &events[idx].start, &events[idx].end)
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].start < events[j].start
	})
	ans := make([]event, n)
	for idx := 0; idx < n; idx++ {
		ans[idx] = event{
			friend: idx,
			start:  -1,
			end:    -1,
		}
	}

	cstart := 0
	cend := 0
	cidx := -1
	for idx := 0; idx < n; idx++ {
		if events[idx].end > cend {
			if cidx != -1 {
				ans[cidx] = event{
					friend: cidx,
					start:  cstart,
					end:    min(cend, events[idx].start),
				}

			}
			cstart = events[idx].start
			cend = events[idx].end
			cidx = events[idx].friend
		}
	}

	if cidx != -1 {
		ans[cidx] = event{
			friend: cidx,
			start:  cstart,
			end:    cend,
		}
	}

	for _, a := range ans {
		write("%d %d\n", a.start, a.end)
	}
}

func read(format string, a ...any) {
	_, _ = fmt.Fscanf(reader, format, a...)
}

func write(format string, a ...any) {
	_, _ = fmt.Fprintf(writer, format, a...)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
