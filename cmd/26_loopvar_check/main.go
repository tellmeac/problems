package main

import (
	"fmt"
)

func main() {
	a := []int{1, 2, 3}
	b := make([]*int, len(a))
	for i, v := range a {
		b[i] = &v
	}

	fmt.Println(*b[0], *b[1])
	// 1 2 printed // "loopvar" thing is not a problem anymore
}
