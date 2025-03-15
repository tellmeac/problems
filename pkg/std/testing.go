package std

import (
	"bytes"
	"context"
	"io"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type (
	SolutionFunc func(io.Reader, io.Writer) error
	ValidateFunc func(t *testing.T, r io.Reader) error
)

type TestBuilder interface {
	WithData(in io.Reader) TestBuilder
	WithValidator(ValidateFunc) TestBuilder
	WithTimeout(time.Duration) TestBuilder
	Run(t *testing.T)
}

type std struct {
	data      io.Reader
	fn        SolutionFunc
	validator ValidateFunc
	timeout   time.Duration
}

// New creates a builder for testing solutions.
func New(fn SolutionFunc) TestBuilder {
	return &std{
		fn: fn,
	}
}

func (s *std) WithValidator(v ValidateFunc) TestBuilder {
	s.validator = v
	return s
}

func (s *std) WithData(r io.Reader) TestBuilder {
	s.data = r
	return s
}

func (s *std) WithTimeout(t time.Duration) TestBuilder {
	s.timeout = t
	return s
}

func (s *std) Run(t *testing.T) {
	ctx, cancel := context.WithTimeout(t.Context(), s.timeout)
	defer cancel()

	var gotResult bytes.Buffer

	doneC := make(chan struct{})
	go func() {
		err := s.fn(s.data, &gotResult)
		if err != nil {
			t.Logf("error solution: %v", err)
			t.FailNow()
		}
		close(doneC)
	}()

	select {
	case <-ctx.Done():
		t.Logf("timeout exceeded: %s", s.timeout)
		t.Fail()
	case <-doneC:
		if err := s.validator(t, &gotResult); err != nil {
			t.Logf("validator error: %s", err)
			t.Fail()
		}
	}
}

func ValidateStrict(want string) ValidateFunc {
	return func(t *testing.T, r io.Reader) error {
		d, err := io.ReadAll(r)
		if err != nil {
			return err
		}

		require.Equal(t, want, string(d))

		return nil
	}
}
