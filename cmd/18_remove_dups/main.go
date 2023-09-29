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
	if n <= 0 {
		return
	}

	var val, prev int

	// first iter
	{
		read("%d\n", &val)
		n--
		prev = val
		write("%d\n", val)
	}

	for ; n > 0; n-- {
		read("%d\n", &val)

		if prev != val {
			write("%d\n", val)
			prev = val
		}
	}
}

func read(format string, a ...any) {
	_, _ = fmt.Fscanf(reader, format, a...)
}

func write(format string, a ...any) {
	_, _ = fmt.Fprintf(writer, format, a...)
}
