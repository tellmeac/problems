package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeCycles("UDRULD"))
	fmt.Println(removeCycles("LDLUR"))
	fmt.Println(removeCycles("URRLLR"))
}

/*
Убрать петли из строки

U - up
R - right
D - down
L - left

"UDRULD" -> ""
"LDLUR" -> "L"
"URRLLR" -> "UR"

Решение:

1. Двигаться вправо и сохранять в таблицу (x,y) -> first idx
2. Двигаться вправо и записывать ответ

*/

type Position struct {
	X, Y int
}

func (p Position) Add(move Position) Position {
	return Position{p.X + move.X, p.Y + move.Y}
}

func removeCycles(moves string) string {
	directions := map[rune]Position{
		'L': {X: -1, Y: 0},
		'R': {X: 1, Y: 0},
		'U': {X: 0, Y: 1},
		'D': {X: 0, Y: -1},
	}

	result := make([]rune, 0, len(moves))

	current := Position{0, 0}

	visited := make(map[Position]int, len(moves))
	visited[current] = -1

	for i, move := range moves {
		current = current.Add(directions[move])

		j, ok := visited[current]
		if !ok {
			visited[current] = i
			result = append(result, move)
			continue
		}

		result = result[:j+1]
	}

	return string(result)
}
