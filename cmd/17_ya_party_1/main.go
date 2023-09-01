package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
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

	ans := 1

	var raw string
	read("%s\n", &raw)
	current, _ := time.Parse("15:04:05", raw)

	for idx := 1; idx < n; idx++ {
		read("%s\n", &raw)
		next, _ := time.Parse("15:04:05", raw)

		if current.After(next) || current == next {
			ans++
		}

		current = next
	}

	write("%d", ans)
}

func read(format string, a ...any) {
	_, _ = fmt.Fscanf(reader, format, a...)
}

func write(format string, a ...any) {
	_, _ = fmt.Fprintf(writer, format, a...)
}
