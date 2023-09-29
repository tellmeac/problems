package main

import "fmt"

// 1) Дана строка из десятичных цифр (длинное число, младшие разряды расположены по младшему индексу).
// Написать код, который умножит это число на число 1 <= n <= 9.

func Multiple(num []int, n int) []int {
	// 123 * 5 = 615
	// [3 2 1] * 5 -> 3 * 5 + 20 * 5 + 100 * 5
	// [5] + 1, [5, 1] + 1, [5, 1, 6] - okay

	ans := num[:]

	remainder := 0

	// [5] + 1, [5, 1] + 1, [5, 1, 6] - okay
	for idx := range num {
		x := ans[idx] * n

		ans[idx] = x%10 + remainder

		remainder = x / 10
	}

	if remainder > 0 {
		ans = append(ans, remainder)
	}

	return ans
}

func main() {
	fmt.Print(Multiple([]int{3, 2, 1}, 5)) // returns 5 1 6 -> because 123 * 5 = 615
}
