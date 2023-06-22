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

	p := 0
	prices := make([]int, 0, n)
	for ; n > 0; n-- {
		_, _ = fmt.Fscanf(reader, "%d", &p)
		prices = append(prices, p)
	}

	if len(prices) == 1 {
		_, _ = fmt.Fprint(writer, "0 0\n")
		return
	}

	const deposit = 1000

	bestBuyDay := 0
	BestCellDay := 0
	minCostDay := 0
	for index := 1; index < len(prices); index++ {
		if prices[BestCellDay]*prices[minCostDay] < prices[bestBuyDay]*prices[index] {
			bestBuyDay = minCostDay
			BestCellDay = index
		}

		if prices[index] < prices[minCostDay] {
			minCostDay = index
		}
	}

	if bestBuyDay == 0 && BestCellDay == 0 {
		_, _ = fmt.Fprint(writer, "0 0\n")
	} else {
		_, _ = fmt.Fprintf(writer, "%d %d\n", bestBuyDay+1, BestCellDay+1)
	}

	_ = writer.Flush()
}

// func max(a, b int) int {
// 	if a < b {
// 		return b
// 	}
// 	return a
// }

// func min(a, b int) int {
// 	if a > b {
// 		return b
// 	}
// 	return a
// }
