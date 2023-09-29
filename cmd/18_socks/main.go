package main

import (
	"bufio"
	"fmt"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer func() {
		_ = writer.Flush()
	}()

	var l, n, m int
	read("%d %d %d\n", &l, &n, &m)

	balance := make([]int, l+1)
	ans := make([]int, l+1)

	var left, right int
	for idx := 0; idx < n; idx++ {
		read("%d %d\n", &left, &right)
		balance[left-1] += 1
		balance[right] -= 1
	}

	now := 0
	for idx := 0; idx < l; idx++ {
		now += balance[idx]
		ans[idx] = now
	}

	var query int
	for idx := 0; idx < m; idx++ {
		read("%d\n", &query)
		query--
		write("%d\n", ans[query])
	}
}

func read(format string, a ...any) {
	_, _ = fmt.Fscanf(reader, format, a...)
}

func write(format string, a ...any) {
	_, _ = fmt.Fprintf(writer, format, a...)
}
