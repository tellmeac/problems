package lcode

import (
	"container/heap"

	"github.com/tellmeac/problems/pkg/pqueue"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeLists(lists []*ListNode) *ListNode {
	/*
		Use max heap this way:
		1. build from each list's first element
		2. pop elem from queue to add into answer
		3. when element is pop add it from one of lists
		4.
	*/

	pq := &pqueue.PriorityQueue{Values: make([]*pqueue.Item, 0, len(lists))}
	for idx := range lists {
		if lists[idx] == nil {
			continue
		}

		top := lists[idx]

		pq.Push(&pqueue.Item{
			Value:    top,
			Priority: top.Val,
		})
	}

	heap.Init(pq)

	dummy := &ListNode{Val: -1}
	p := dummy

	for pq.Len() > 0 {
		heapItem := heap.Pop(pq).(*pqueue.Item)
		node := heapItem.Value.(*ListNode)

		p.Next = node
		p = p.Next

		if node.Next != nil {
			heap.Push(pq, &pqueue.Item{
				Value:    node.Next,
				Priority: node.Next.Val,
			})
		}
	}

	return dummy.Next
}
