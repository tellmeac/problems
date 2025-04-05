package lcode

func findAnagrams(text string, subStr string) (result []int) {
	// fast path
	if len(subStr) > len(text) {
		return nil
	}

	weakHash := func(s string) int64 {
		r := int64(0)
		for _, ch := range s {
			r += int64(ch) * int64(ch) * int64(ch) * int64(ch)
		}
		return r
	}

	// removes idx-1 from hash value and adds idx+len(subStr) to hash.
	updateHash := func(val int64, idx int) int64 {
		ch := text[idx-1]
		val -= int64(ch) * int64(ch) * int64(ch) * int64(ch)
		ch = text[idx+len(subStr)-1]
		val += int64(ch) * int64(ch) * int64(ch) * int64(ch)

		return val
	}

	reference := weakHash(subStr)
	candidate := weakHash(text[:len(subStr)])
	if reference == candidate {
		result = append(result, 0)
	}

	// invariants:
	// (i) is the start of possible anagram
	// (i + len(subStr) - 1) is the end of possible anagram
	// updateHash called with i > 0
	for i := 1; i <= len(text)-len(subStr); i++ {
		candidate = updateHash(candidate, i)
		if reference == candidate {
			result = append(result, i)
		}
	}

	return result
}
