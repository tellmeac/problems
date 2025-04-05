package pp

import (
	"sync"
	"testing"
	"time"
)

type Service struct {
	Values []any
}

func (s Service) Set(i int, val any) {
	s.Values[i] = val
}

func (s Service) Get(i int) any {
	return s.Values[i]
}

// go test not failing, but
// go test -race must fail
func TestRace(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	var wg sync.WaitGroup
	wg.Add(2)

	s := Service{Values: make([]any, 2)}

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ticker.C:
				return
			default:
				s.Set(0, 1)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ticker.C:
				return
			default:
				_ = s.Get(0)
			}
		}
	}()

	wg.Wait()
}
