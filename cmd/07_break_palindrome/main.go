package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	var s string

	_, _ = fmt.Fscanf(reader, "%s\n", &s)

	s = strings.TrimSpace(s)
	if len(s) <= 1 {
		return
	}

	// "a" is 97
	pal := []byte(s)
	turned := false
	for ind, c := range pal {
		if c > 97 {
			if len(pal)%2 != 0 && ind == len(pal)/2 {
				continue
			}
			pal[ind] = 97
			turned = true
			break
		}
	}

	if !turned {
		pal[len(pal)-1] += 1
	}

	_, _ = fmt.Fprintf(writer, "%s", string(pal))

	_ = writer.Flush()
}
