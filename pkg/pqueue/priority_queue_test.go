package pqueue

import (
	"container/heap"
	"fmt"
	"strings"
)

func ExamplePriorityQueue() {
	data := map[string]int{
		"Alexander": 2,
		"Andrey":    4,
		"Sergey":    8,
	}

	pq := &PriorityQueue{
		Values: make([]*Item, 0, len(data)),
	}
	for value, priority := range data {
		pq.Push(&Item{
			Value:    value,
			Priority: -priority, // NOTE: negative priority!
		})
	}

	heap.Init(pq)

	result := make([]string, 0, len(data))
	for range data {
		item := heap.Pop(pq).(*Item)
		result = append(result, item.Value.(string))
	}

	fmt.Print(strings.Join(result, " "))
	// Output: Sergey Andrey Alexander
}
