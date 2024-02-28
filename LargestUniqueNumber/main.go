package LargestUniqueNumber

func largestUniqueNumber(nums []int) int {
	var mapped = make(map[int]int)
	for _, num := range nums {
		mapped[num]++
	}
	var biggest = -1
	for key, val := range mapped {
		if val == 1 && biggest < key {
			biggest = key
		}
	}
	return biggest
}
