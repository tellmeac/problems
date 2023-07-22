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

	var n int
	read("%d\n", &n)

	users := make([]int, n)
	for idx := 0; idx < n; idx++ {
		read("%d", &users[idx])
	}

	sort.Slice(users, func(i, j int) bool {
		return users[i] < users[j]
	})

	l := 0
	r := 0
	ans := 0
	for idx := 0; idx < n; idx++ {
		for l < n && users[l] <= (users[idx]/2)+7 {
			l++
		}
		for r < n && users[r] <= users[idx] {
			r++
		}

		if r > l+1 {
			ans += r - l - 1
		}
	}
	write("%d\n", ans)
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
