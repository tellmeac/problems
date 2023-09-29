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

	var a, b string
	read("%s\n%s\n", &a, &b)

	statA := make([]int, 26)
	for _, val := range []byte(a) {
		statA[val-'a']++
	}

	for _, val := range []byte(b) {
		statA[val-'a']--
	}

	// check

	for _, v := range statA {
		if v != 0 {
			write("0\n")
			return
		}
	}

	write("1\n")
}

func read(format string, a ...any) {
	_, _ = fmt.Fscanf(reader, format, a...)
}

func write(format string, a ...any) {
	_, _ = fmt.Fprintf(writer, format, a...)
}
