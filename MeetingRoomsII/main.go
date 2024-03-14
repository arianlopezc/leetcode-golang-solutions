package MeetingRoomsII

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h MinHeap) Peek() int {
	if h.Len() > 0 {
		return h[0]
	} else {
		return -1
	}
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minMeetingRooms(intervals [][]int) int {
	var currRooms, totalRooms = 0, 0
	var startTimes, endTimes = &MinHeap{}, &MinHeap{}
	heap.Init(startTimes)
	heap.Init(endTimes)
	for _, interval := range intervals {
		heap.Push(startTimes, interval[0])
		heap.Push(endTimes, interval[1])
	}
	for startTimes.Len() > 0 {
		for endTimes.Len() > 0 && startTimes.Peek() >= endTimes.Peek() {
			heap.Pop(endTimes)
			currRooms--
		}
		currRooms++
		if currRooms > totalRooms {
			totalRooms = currRooms
		}
		heap.Pop(startTimes)
	}
	return totalRooms
}
