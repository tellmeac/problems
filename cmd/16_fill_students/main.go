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

func main() {
	defer func() {
		_ = writer.Flush()
	}()

	var n int // students groups
	read("%d\n", &n)

	groups := make([]int, n)
	for idx := 0; idx < len(groups); idx++ {
		read("%d", &groups[idx])
	}
	read("\n")

	var k int // free rooms
	read("%d\n", &k)

	rooms := make([]int, k)
	for idx := 0; idx < len(rooms); idx++ {
		read("%d", &rooms[idx])
	}

	sort.Slice(groups, func(i, j int) bool {
		return groups[i] > groups[j]
	})

	sort.Slice(rooms, func(i, j int) bool {
		return rooms[i] > rooms[j]
	})

	ans := 0

	// merging rooms and groups
	idx := 0 // group idx
	jdx := 0 // room idx
	for idx < len(groups) && jdx < len(rooms) {
		g := groups[idx]
		r := rooms[jdx]

		if g <= r {
			ans++
			idx++
			jdx++
		} else {
			idx++
		}
	}

	write("%d", ans)
}

func read(format string, a ...any) {
	_, _ = fmt.Fscanf(reader, format, a...)
}

func write(format string, a ...any) {
	_, _ = fmt.Fprintf(writer, format, a...)
}
