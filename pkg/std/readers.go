package std

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func ReadIntegers(in io.Reader) ([]int, error) {
	var s []int

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		val, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, fmt.Errorf("parse int: %w", err)
		}

		s = append(s, val)
	}

	return s, nil
}
