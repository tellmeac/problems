package pqueue

import (
	"container/heap"
)

var (
	_ heap.Interface = &PriorityQueue{}
)

type Item struct {
	Value    interface{}
	Priority int

	index int
}

type PriorityQueue struct {
	Values []*Item
}

func (pq *PriorityQueue) Len() int {
	return len(pq.Values)
}

func (pq *PriorityQueue) Less(i, j int) bool {
	return pq.Values[i].Priority < pq.Values[j].Priority
}

func (pq *PriorityQueue) Swap(i, j int) {
	pq.Values[i], pq.Values[j] = pq.Values[j], pq.Values[i]

	pq.Values[i].index = i
	pq.Values[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	item, ok := x.(*Item)
	if !ok {
		panic("expected pushed item to be an *Item")
	}

	item.index = len(pq.Values)

	pq.Values = append(pq.Values, item)
}

func (pq *PriorityQueue) Pop() any {
	old := pq.Values
	n := len(old)

	item := old[n-1]
	old[n-1] = nil

	item.index = -1

	pq.Values = old[0 : n-1]

	return item
}
