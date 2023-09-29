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

var n int

func main() {
	defer func() {
		_ = writer.Flush()
	}()

	read("%d\n", &n)

	ans := make([]string, 0)

	generator(0, 0, make([]byte, 2*n), &ans)

	for _, a := range ans {
		write("%s\n", a)
	}
}

func generator(l, r int, gen []byte, out *[]string) {
	if l+r == 2*n {
		*out = append(*out, string(gen))
		return
	}

	if l < n {
		gen[l+r] = '('
		generator(l+1, r, gen, out)
	}

	if r < l {
		gen[l+r] = ')'
		generator(l, r+1, gen, out)
	}
}

func read(format string, a ...any) {
	_, _ = fmt.Fscanf(reader, format, a...)
}

func write(format string, a ...any) {
	_, _ = fmt.Fprintf(writer, format, a...)
}
