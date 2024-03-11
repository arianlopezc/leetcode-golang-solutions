package MissingElementInSortedArray

func totalMissingNumbers(nums []int, index int) int {
	return nums[index] - nums[0] - index
}

func missingElement(nums []int, k int) int {
	if totalMissingNumbers(nums, len(nums)-1) < k {
		return nums[len(nums)-1] + k - totalMissingNumbers(nums, len(nums)-1)
	}
	var left = 0
	var right = len(nums) - 1
	for left != right {
		var pivot = left + (right-left)/2
		if totalMissingNumbers(nums, pivot) < k {
			left = pivot + 1
		} else {
			right = pivot
		}
	}
	return nums[left-1] + k - totalMissingNumbers(nums, left-1)
}
