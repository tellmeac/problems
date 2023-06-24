package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

type DictStore map[string]struct{}

func (store DictStore) Replace(word string) (string, bool) {
	for index := len(word); index >= 0; index-- {
		dw := word[:len(word)-index]
		if _, ok := store[dw]; ok {
			return dw, true
		}
	}

	return "", false
}

func main() {
	defer func() {
		_ = writer.Flush()
	}()

	scanner := bufio.NewScanner(reader)

	scanner.Scan()
	var dictRaw = strings.TrimSpace(scanner.Text())
	store := make(DictStore, 0)
	for _, w := range strings.Split(dictRaw, " ") {
		store[w] = struct{}{}
	}

	scanner.Scan()
	var text = strings.TrimSpace(scanner.Text())
	answer := make([]string, 0)
	for _, w := range strings.Split(text, " ") { // TODO: proper split required: "a  b a bb    b".
		r, ok := store.Replace(w)
		if !ok {
			answer = append(answer, w)
		} else {
			answer = append(answer, r)
		}
	}

	write("%s\n", strings.Join(answer, " "))
}

func write(format string, a ...any) {
	_, _ = fmt.Fprintf(writer, format, a...)
}
