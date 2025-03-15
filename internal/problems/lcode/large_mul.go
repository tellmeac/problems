package lcode

// Multiple multiplies a large number with n.
func Multiple(num []uint64, n uint64) []uint64 {
	// 123 * 5 = 615
	// [3 2 1] * 5 -> 3 * 5 + 20 * 5 + 100 * 5
	// [5] + 1, [5, 1] + 1, [5, 1, 6] - okay

	result := make([]uint64, 0, len(num))
	var remainder uint64

	// [5] + 1, [5, 1] + 1, [5, 1, 6] - okay
	for idx := range num {
		current := result[idx] * n

		result[idx] = current%10 + remainder

		// update remainder
		remainder = current / 10
	}

	if remainder > 0 {
		result = append(result, remainder)
	}

	return result
}
