package TheKWeakestRowsInAMatrix

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
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

func totalSoldiers(row []int) int {
	var left, right = 0, len(row) - 1
	for left <= right {
		var middle = (right-left)/2 + left
		var middleNext = middle + 1
		var curr = row[middle]
		var nxt int
		if len(row) > middleNext {
			nxt = row[middleNext]
		}
		if curr == 1 && nxt == 0 {
			return middle + 1
		} else if curr == 1 && nxt == 1 {
			left = middle + 1
		} else if curr == 0 && nxt == 0 {
			if middle == 0 {
				return 0
			}
			right = middle - 1
		}
	}
	return 0
}

func kWeakestRows(mat [][]int, k int) []int {
	var weakness = make(map[int]*IntHeap)
	var heapWeak = &IntHeap{}
	heap.Init(heapWeak)
	for index, row := range mat {
		var soldiersInRow = totalSoldiers(row)
		if rows, ok := weakness[soldiersInRow]; ok {
			heap.Push(rows, index)
		} else {
			h := &IntHeap{index}
			heap.Init(h)
			weakness[soldiersInRow] = h
			heap.Push(heapWeak, soldiersInRow)
		}
	}
	var result []int
	for k > 0 && heapWeak.Len() > 0 {
		var weak = heap.Pop(heapWeak).(int)
		for weakness[weak].Len() > 0 && k > 0 {
			var row = heap.Pop(weakness[weak]).(int)
			result = append(result, row)
			k--
		}
	}
	return result
}
