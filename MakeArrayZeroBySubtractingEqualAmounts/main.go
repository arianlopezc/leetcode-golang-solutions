package MakeArrayZeroBySubtractingEqualAmounts

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Peek() int {
	if len(*h) > 0 {
		return (*h)[0]
	}
	return -1
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minimumOperations(nums []int) int {
	var heapQueue = &IntHeap{}
	heap.Init(heapQueue)
	for _, val := range nums {
		if val != 0 {
			heap.Push(heapQueue, val)
		}
	}
	var totalOps int
	for heapQueue.Len() > 0 {
		var front int
		for heapQueue.Len() > 0 && heapQueue.Peek() == 0 {
			heap.Pop(heapQueue)
		}
		if heapQueue.Len() > 0 {
			front = heap.Pop(heapQueue).(int)
			totalOps++
			for index := 0; index < heapQueue.Len(); index++ {
				(*heapQueue)[index] -= front
				if (*heapQueue)[index] <= 0 {
					(*heapQueue)[index] = 0
				}
			}
		}
	}
	return totalOps
}
