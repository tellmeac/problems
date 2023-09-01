package main

import (
	"errors"
	"fmt"
	"sync"
)

type Value interface {
	F() error
}

func Scheduler(s ...Value) []error {
	errC := make(chan error, 1)

	res := make([]error, 0)
	go func() {
		for err := range errC {
			res = append(res, err)
		}
	}()

	wg := &sync.WaitGroup{}
	wg.Add(len(s))
	for idx := range s {
		idx := idx
		go func() {
			err := s[idx].F()
			if err != nil {
				errC <- err
			}
			wg.Done()
		}()
	}

	wg.Wait()
	close(errC)

	return res
}

type TestValue struct {
	Err error
}

func (t TestValue) F() error {
	return t.Err
}

func main() {
	a := TestValue{Err: errors.New("world")}
	b := TestValue{Err: errors.New("hello")}

	res := Scheduler(a, a, b, b, a)

	fmt.Printf("%+v\n", res)
}
