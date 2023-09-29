package main

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

var (
	reader = bufio.NewReader(os.Stdin)
	writer = bufio.NewWriter(os.Stdout)
)

func main() {
	defer func() {
		_ = writer.Flush()
	}()

	var targetName string
	read("%s\n", &targetName)

	var n int
	read("%d\n", &n)

	var targetIdx int
	hierarchyRaw := make([]string, 0, n)
	scanner := bufio.NewScanner(reader)
	for idx := 0; idx < n; idx++ {
		scanner.Scan()
		v := scanner.Text()
		if name(v) == targetName {
			targetIdx = idx
		}

		hierarchyRaw = append(hierarchyRaw, v)
	}

	if level(hierarchyRaw[targetIdx]) == 0 {
		write("%s\n", path.Join("/", targetName))
		return
	}

	pathParts := make([]string, 0, 2)
	pathParts = append(pathParts, targetName)
	currentLvl := level(hierarchyRaw[targetIdx])
	cursorIdx := targetIdx
	for cursorIdx >= 0 {
		if !isDir(hierarchyRaw[cursorIdx]) {
			cursorIdx--
			continue
		}

		lvl := level(hierarchyRaw[cursorIdx])
		if lvl+1 == currentLvl {
			pathParts = append(pathParts, name(hierarchyRaw[cursorIdx]))
			currentLvl--
		}
		cursorIdx--
	}

	pathParts = append(pathParts, "/")

	for i, j := 0, len(pathParts)-1; i < j; i, j = i+1, j-1 {
		pathParts[i], pathParts[j] = pathParts[j], pathParts[i]
	}

	write("%s\n", path.Join(pathParts...))
}

func read(format string, a ...any) {
	_, _ = fmt.Fscanf(reader, format, a...)
}

func write(format string, a ...any) {
	_, _ = fmt.Fprintf(writer, format, a...)
}

func name(raw string) string {
	return strings.TrimSpace(raw)
}

func level(raw string) int {
	return strings.Count(raw, " ")
}

func isDir(name string) bool {
	return !strings.Contains(name, ".")
}
