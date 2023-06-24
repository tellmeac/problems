package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

const colorsToComplete = 3

type StickInfo struct {
	Rings map[byte]struct{}
}

func (s *StickInfo) Add(color byte) {
	s.Rings[color] = struct{}{}
}

func (s *StickInfo) IsCompleted() bool {
	var r int

	for range s.Rings {
		r++
	}

	return r >= colorsToComplete
}

func main() {
	defer func() {
		_ = writer.Flush()
	}()

	var s string
	read("%s", &s)

	if len(s)%2 != 0 {
		panic("expected to have size % 2 == 0")
	}

	data := []byte(s)
	sticks := make(map[int]*StickInfo, 10) // 10 by definition
	for index := 0; index < len(s)-1; index += 2 {
		color := data[index]
		stickId, _ := strconv.Atoi(string(data[index+1]))

		info, ok := sticks[stickId]
		if !ok {
			info = &StickInfo{Rings: map[byte]struct{}{}}
			info.Add(color)
			sticks[stickId] = info
		} else {
			sticks[stickId].Add(color)
		}
	}

	answer := 0
	for _, info := range sticks {
		if info.IsCompleted() {
			answer++
		}
	}

	write("%d", answer)
}

func read(format string, a ...any) {
	_, _ = fmt.Fscanf(reader, format, a...)
}

func write(format string, a ...any) {
	_, _ = fmt.Fprintf(writer, format, a...)
}
