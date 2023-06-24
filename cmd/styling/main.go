package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"sort"
	"time"
)

var ignoredSolutions = map[string]struct{}{
	"example": {},
	"styling": {},
}

const (
	solutionsDir = "./cmd"
)

type definition struct {
	name    string
	modTime time.Time
}

func main() {
	entries, err := os.ReadDir(solutionsDir)
	if err != nil {
		log.Fatalf("Failed to read dir: %s\n", err)
		return
	}

	defs := make([]definition, 0, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		if _, isIgnored := ignoredSolutions[entry.Name()]; isIgnored {
			continue
		}

		info, err := entry.Info()
		if err != nil {
			log.Fatalf("failed to get entry info: %s", err)
		}
		defs = append(defs, definition{
			name:    entry.Name(),
			modTime: info.ModTime(),
		})
	}

	sort.Slice(defs, func(i, j int) bool {
		return defs[i].modTime.Before(defs[j].modTime)
	})

	expr, err := regexp.Compile("\\d_*")
	if err != nil {
		log.Fatalf("invalid expression: %s", err)
	}

	for idx, def := range defs {
		oldPath := path.Join(solutionsDir, def.name)
		if expr.MatchString(oldPath) {
			log.Printf("dir already numbered: %q", oldPath)
			continue
		}

		newPath := path.Join(solutionsDir, fmt.Sprintf("%02d_%s", idx, def.name))
		if err := os.Rename(oldPath, newPath); err != nil {
			log.Printf("Failed to rename dir: %s", err)
		}
	}
}
