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
	var n int

	_, _ = fmt.Fscanf(reader, "%d\n", &n)

	for ; n > 0; n-- {
		_, _ = fmt.Fprintf(writer, "%d\n", n)
	}

	_ = writer.Flush()
}
