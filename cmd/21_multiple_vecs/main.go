package main

import (
	"fmt"
)

type Zip struct {
	Value int
	Count int
}

//func iterator(s []Zip) func() (int, bool) {
//	idx := 0
//
//	var cb func() (int, bool)
//	cb = func() (int, bool) {
//		if idx >= len(s) {
//			return 0, false
//		}
//
//		if s[idx].Count > 0 {
//			s[idx].Count--
//
//			return s[idx].Value, true
//		}
//
//		idx++
//		return cb()
//	}
//
//	return cb
//}
//
//// scalar multiple
//// zip vector [(1, 2), (10, 1)] -> is [1 1 10].
//func scalar(a, b []Zip) int {
//	// know your limits, valid vectors contains N values each
//
//	aIter := iterator(a)
//	bIter := iterator(b)
//
//	ans := 0
//
//	v1, ok1 := aIter()
//	v2, ok2 := bIter()
//
//	for ok1 && ok2 {
//		ans += v1 * v2
//
//		v1, ok1 = aIter()
//		v2, ok2 = bIter()
//	}
//
//	return ans
//}

func Scalar(a, b []Zip) int {
	// a - [1 100]
	// b - [2 50] [3 50]

	i, j := 0, 0
	f := func(x, y *Zip) int {
		if x.Count <= 0 {
			i++
			return 0
		}

		if y.Count <= 0 {
			j++
			return 0
		}

		m := min(x.Count, y.Count)

		x.Count -= m
		y.Count -= m

		return x.Value * y.Value * m
	}

	ans := 0
	for i < len(a) && j < len(b) {
		ans += f(&a[i], &b[j])
	}

	return ans
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	// Пример [(1,3)], [(1,2), (10,1)] -> 12

	fmt.Print(Scalar([]Zip{{1, 3}}, []Zip{{1, 2}, {10, 1}}))
}
