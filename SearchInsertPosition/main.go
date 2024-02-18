package SearchInsertPosition

func searchInsert(nums []int, target int) int {
	var l, r = 0, len(nums) - 1
	for l <= r {
		index := l + (r-l)/2
		if nums[index] == target {
			return index
		} else if nums[index] < target {
			l = index + 1
		} else {
			r = index - 1
		}
	}
	return l
}
