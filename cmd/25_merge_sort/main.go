package main

import (
	"fmt"
)

func merge(a, b []int) []int {
	s := make([]int, len(a)+len(b))

	var (
		left     = 0
		right    = 0
		writeIdx = 0
	)
	for left < len(a) && right < len(b) {
		if a[left] < b[right] {
			s[writeIdx] = a[left]
			left++
			writeIdx++
			continue
		}
		s[writeIdx] = b[right]
		right++
		writeIdx++
	}

	for left < len(a) {
		s[writeIdx] = a[left]
		writeIdx++
		left++
	}

	for right < len(b) {
		s[writeIdx] = b[right]
		writeIdx++
		right++
	}

	if writeIdx != len(s) {
		panic("invalid solution")
	}

	return s
}

func main() {
	l := []int{1, 1, 2, 2, 3, 4}
	r := []int{1, 3, 3, 4, 5, 50}

	fmt.Print(merge(l, r))
}
