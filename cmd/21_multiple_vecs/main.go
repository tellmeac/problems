package main

type Zip struct {
	Value int
	Count int
}

func Scalar(a, b []Zip) int {
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
