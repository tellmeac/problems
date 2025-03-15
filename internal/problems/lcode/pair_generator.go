package lcode

func PairGenerator(n int) []string {
	const (
		left  = '('
		right = ')'
	)

	var result []string

	var generator func(l, r int, gen []byte)
	generator = func(l, r int, current []byte) {
		// write completed
		if l+r == 2*n {
			result = append(result, string(current))
			return
		}

		// move left
		if l < n {
			current[l+r] = left
			generator(l+1, r, current)
		}

		// move right
		if r < l {
			current[l+r] = right
			generator(l, r+1, current)
		}
	}

	generator(0, 0, make([]byte, 2*n))

	return result
}
