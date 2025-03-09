package main

import (
	"fmt"
	"sort"
)

func findGood(provided, requests []int) int {
	sort.Ints(provided)

	closest := func(val int) int {
		s := provided

		low, high := 0, len(s)-1

		for low <= high {
			mid := low + (high-low)/2

			if s[mid] == val {
				return val
			} else if s[mid] < val {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

		a := val - s[low]
		if a < 0 {
			a = -a
		}
		b := val - s[high]
		if b < 0 {
			b = -b
		}

		return min(a, b)
	}

	result := 0
	for _, r := range requests {
		good := r - closest(r)
		if good < 0 {
			good = -good
		}

		result += good
	}

	return result
}

func main() {
	fmt.Println(findGood([]int{8, 3, 5}, []int{5, 14, 12, 44, 55}))
}
