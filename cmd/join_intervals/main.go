package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	//
}

func compress(s []int) string {
	sort.Slice(s, func(i, j int) bool {
		return s[i] < s[j]
	})

	ans := make([]string, 0)

	fromIdx := 0
	isIntervalCompleted := true
	for idx := 0; idx < len(s)-1; idx++ {
		if s[idx+1]-s[idx] == 1 {
			isIntervalCompleted = false
			continue
		}

		// Ниже нужно записать ответ и переместить начало следующего интервала для ответа

		if fromIdx != idx {
			ans = append(ans, fmt.Sprintf("%d-%d", s[fromIdx], s[idx]))
		} else {
			ans = append(ans, fmt.Sprintf("%d", s[fromIdx]))
		}

		isIntervalCompleted = true
		fromIdx = idx + 1
	}

	if isIntervalCompleted {
		// запишем последнее число
		ans = append(ans, fmt.Sprintf("%d", s[fromIdx]))
	} else {
		// запишем последний интервал
		ans = append(ans, fmt.Sprintf("%d-%d", s[fromIdx], s[len(s)-1]))
	}

	// 1 2 3 4 7 8
	// 1 2 3 4 7 8 10

	return strings.Join(ans, ",")
}
