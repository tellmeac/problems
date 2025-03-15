package problems_test

import (
	"strings"
	"testing"
	"time"

	"github.com/tellmeac/problems/internal/problems"
	"github.com/tellmeac/problems/pkg/std"
)

func TestExampleSolution(t *testing.T) {
	fn := std.SolutionFunc(problems.ExampleSolution)

	std.New(fn).
		WithData(strings.NewReader("1 2 3")).
		WithValidator(std.ValidateStrict("6")).
		WithTimeout(time.Second).
		Run(t)
}
