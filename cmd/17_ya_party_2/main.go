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

	valid := true
	var from, to string
	for reqIdx := 0; reqIdx < n; reqIdx++ {
		read("%s\n", &from)
		read("%s\n", &to)

		ds := make(map[byte]byte, 26)
		dt := make(map[byte]byte, 26)
		s := []byte(from)
		t := []byte(to)

		for idx := 0; idx < len(s); idx++ {
			a, ok := ds[s[idx]]
			if !ok {
				ds[s[idx]] = t[idx]
			} else {
				if a != t[idx] {
					write("NO\n")
					valid = false
					break
				}
			}

			b, ok := dt[s[idx]]
			if !ok {
				dt[t[idx]] = s[idx]
			} else {
				if b != s[idx] {
					write("NO\n")
					valid = false
					break
				}
			}
		}
		if valid {
			write("YES\n")
		}
	}
}

func read(format string, a ...any) {
	_, _ = fmt.Fscanf(reader, format, a...)
}

func write(format string, a ...any) {
	_, _ = fmt.Fprintf(writer, format, a...)
}
