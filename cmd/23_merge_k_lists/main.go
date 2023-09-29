package main

import "container/heap"

var (
	_ heap.Interface = &PriorityQueue{}
)

type Item struct {
	Value    *ListNode
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
	return pq.Values[i].Priority < pq.Values[j].Priority // change logic to make it MIN / MAX Priority queue
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

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	/*
		Use max heap this way:
		1. build from each list's first element
		2. pop elem from queue to add into answer
		3. when element is pop add it from one of lists
		4.
	*/

	pq := &PriorityQueue{Values: make([]*Item, 0, len(lists))}
	for idx := range lists {
		if lists[idx] == nil {
			continue
		}

		top := lists[idx]

		pq.Push(&Item{
			Value:    top,
			Priority: top.Val,
		})
	}

	heap.Init(pq)

	head := &ListNode{Val: -1}
	tail := head

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)

		tail.Next = item.Value
		tail = tail.Next

		if item.Value.Next != nil {
			heap.Push(pq, &Item{
				Value:    item.Value.Next,
				Priority: item.Value.Next.Val,
			})
		}
	}

	return head.Next
}
