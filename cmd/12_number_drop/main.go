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

	var v, answer int
	count := make(map[int]int, 1e2)
	for i := 0; i < n; i++ {
		read("%d", &v)
		count[v]++

		c, ok := count[v]
		if !ok {
			continue
		}

		if c > n/2 {
			answer = v
		}
	}

	write("%d", answer)
}

func read(format string, a ...any) {
	_, _ = fmt.Fscanf(reader, format, a...)
}

func write(format string, a ...any) {
	_, _ = fmt.Fprintf(writer, format, a...)
}
