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

	var n int
	read("%d\n", &n)

	current := 0
	ans := 0

	var v int
	for ; n > 0; n-- {
		read("%d\n", &v)

		if v == 0 {
			ans = max(ans, current)
			current = 0
		} else {
			current++
		}
	}

	write("%d\n", max(ans, current))
}

func read(format string, a ...any) {
	_, _ = fmt.Fscanf(reader, format, a...)
}

func write(format string, a ...any) {
	_, _ = fmt.Fprintf(writer, format, a...)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
