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

	nums := make(map[int]int, 0)
	x := 0
	for i := 0; i < n; i++ {
		read("%d", &x)
		nums[x]++
	}

	ans := 1
	for num, primaryCount := range nums {
		ans = max(ans, primaryCount)

		if secondCount, ok := nums[num+1]; ok {
			ans = max(ans, secondCount+primaryCount)
		}

		if secondCount, ok := nums[num-1]; ok {
			ans = max(ans, secondCount+primaryCount)
		}
	}

	write("%d\n", n-ans)
}

func read(format string, a ...any) {
	_, _ = fmt.Fscanf(reader, format, a...)
}

func write(format string, a ...any) {
	_, _ = fmt.Fprintf(writer, format, a...)
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
