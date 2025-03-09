package main

func main() {
	print(minimumRecolors("WBBWWBBWBW", 7))
}

/*
	Input: blocks = "WBBWWBBWBW", k = 7
	Input: blocks = "0110011010", k = 7
*/

func minimumRecolors(blocks string, k int) int {
	cb := 0     // current blacks
	cw := 0     // current whites
	result := 0 // result

	i := 0
	for i < k {
		if blocks[i] == 'W' {
			cw++
		} else {
			cb++
		}
		i++
	}

	result = cw

	for i = k; i < len(blocks); i++ {
		if blocks[i-k] == 'W' {
			cw--
		} else {
			cb--
		}

		if blocks[i] == 'W' {
			cw++
		} else {
			cb++
		}

		result = min(result, cw)
	}

	return result
}
