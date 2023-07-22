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

	var n, k int
	read("%d %d\n", &n, &k)

	counts := make([]int, k)
	for idx := 0; idx < k; idx++ {
		read("%d\n", &counts[idx])
	}
	sort.Slice(counts, func(i, j int) bool {
		return counts[i] > counts[j]
	})

	ok := func(m int) bool {
		in := 0
		for idx := 0; idx < k; idx++ {
			in += counts[idx] / m
		}
		return in >= n
	}

	l := 0
	r := int(2 * 10e9)
	for l < r {
		m := (l + r + 1) / 2
		if ok(m) {
			l = m
		} else {
			r = m - 1
		}
	}

	write("%d\n", l)

	usedLamps := 0
	for idx := 0; idx < k; idx++ {
		for jdx := 0; jdx < min(counts[idx]/l, n-usedLamps); jdx++ {
			write("%d\n", idx+1)
			usedLamps++
		}
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
