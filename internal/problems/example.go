package problems

import (
	"bufio"
	"fmt"
	"io"
	"strconv"

	"github.com/tellmeac/problems/pkg/std"
)

// ExampleSolution test solution of summing of []int values
func ExampleSolution(in io.Reader, w io.Writer) error {
	s, err := std.ReadIntegers(in)
	if err != nil {
		return err
	}

	var result int
	for _, v := range s {
		result += v
	}

	buff := bufio.NewWriter(w)
	defer func() { _ = buff.Flush() }()

	if _, err := buff.WriteString(strconv.Itoa(result)); err != nil {
		return fmt.Errorf("write result: %w", err)
	}

	return nil
}
