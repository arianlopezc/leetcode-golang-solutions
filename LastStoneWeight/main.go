package LastStoneWeight

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func abs(one int, two int) int {
	var diff = one - two
	if diff < 0 {
		return diff * -1
	}
	return diff
}

func lastStoneWeight(stones []int) int {
	var pQueue = &IntHeap{}
	for _, stone := range stones {
		heap.Push(pQueue, stone)
	}
	heap.Init(pQueue)
	for pQueue.Len() > 1 {
		var firstStone = heap.Pop(pQueue).(int)
		var secondStone = heap.Pop(pQueue).(int)
		if firstStone == secondStone {
			continue
		} else {
			var newRock = abs(firstStone, secondStone)
			heap.Push(pQueue, newRock)
		}
	}
	if pQueue.Len() == 0 {
		return 0
	}
	return heap.Pop(pQueue).(int)
}
