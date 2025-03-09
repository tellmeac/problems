package main

import (
	"cmp"
	"container/heap"
	"fmt"
	"strings"
)

type PQ[V comparable, P cmp.Ordered] struct {
	items []Item[V, P]
}

type Item[V comparable, P cmp.Ordered] struct {
	Value    V
	Priority P
	Index    int
}

func (pq *PQ[V, P]) Len() int { return len(pq.items) }

func (pq *PQ[V, P]) Less(i, j int) bool {
	return pq.items[i].Priority < pq.items[j].Priority
}

func (pq *PQ[V, P]) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.items[i].Index = i
	pq.items[j].Index = j
}

func (pq *PQ[V, P]) Push(val any) {
	n := len(pq.items)
	item := val.(Item[V, P])
	item.Index = n

	pq.items = append(pq.items, item)
}

func (pq *PQ[V, P]) Pop() any {
	var val any
	val, pq.items = pq.items[pq.Len()-1], pq.items[:pq.Len()-1]
	return val
}

func main() {
	h := &PQ[string, int]{}
	h.items = []Item[string, int]{
		{
			Value:    "hello ",
			Priority: 1,
			Index:    0,
		},
		{
			Value:    ", my friends",
			Priority: 3,
			Index:    0,
		},
		{
			Value:    "world",
			Priority: 2,
			Index:    0,
		},
		{
			Value:    "oh, ",
			Priority: 0,
			Index:    0,
		},
	}
	heap.Init(h)

	s := strings.Builder{}
	for h.Len() > 0 {
		item := heap.Pop(h).(Item[string, int])
		s.WriteString(item.Value)
	}

	fmt.Println(s.String())
}
