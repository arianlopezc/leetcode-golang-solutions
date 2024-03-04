package MaximumCountOfPositiveIntegerAndNegativeInteger

func maximumCount(nums []int) int {
	var left, right = 0, len(nums) - 1
	for left <= right {
		middle := (right-left)/2 + left
		var curr = nums[middle]
		var next int
		if middle+1 < len(nums) {
			next = nums[middle+1]
		}
		if curr < 0 && next >= 0 {
			left = middle
			right = left - 1
		} else if curr >= 0 && next >= 0 {
			if middle == 0 {
				left = middle - 1
				right = left - 1
			} else {
				right = middle - 1
			}
		} else {
			left = middle + 1
		}
	}
	right = max(left, 0)
	for right < len(nums) && nums[right] <= 0 {
		right++
	}
	return max(left+1, len(nums)-right)
}
