package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	var n int

	_, _ = fmt.Fscanf(reader, "%d\n", &n)

	// a = 1, b = -1, c = -2*n
	// D = b*b - 4ac
	d := float64(1 + 8*n)

	negative := (1 - math.Sqrt(d)) / 2
	negative *= -1

	_, _ = fmt.Fprintf(writer, "%d", int(math.Floor(negative)))

	_ = writer.Flush()
}
