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

func read(format string, a ...any) {
	_, _ = fmt.Fscanf(reader, format, a...)
}

func write(format string, a ...any) {
	_, _ = fmt.Fprintf(writer, format, a...)
}

func main() {
	defer func() {
		_ = writer.Flush()
	}()

	var n, k int

	read("%d %d\n", &n, &k)

	answer := false
	lastIndex := make(map[int]int, 1e3)
	x := 0
	for index := 0; index < n; index++ {
		read("%d", &x)

		last, ok := lastIndex[x]
		if !ok {
			lastIndex[x] = index
			continue
		}

		if index-last <= k {
			answer = true
		}
		lastIndex[x] = index
	}

	if answer {
		write("YES\n")
	} else {
		write("NO\n")
	}
}
