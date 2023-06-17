package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	var p string
	_, _ = fmt.Fscanf(reader, "%s\n", &p)
	_, _ = fmt.Fprintf(writer, "%s", path.Clean(p))
	_ = writer.Flush()
}
