package MonotonicArray

import "sort"

func isMonotonic(nums []int) bool {
	var asc bool
	var index int
	for index < len(nums)-1 && nums[index] == nums[index+1] {
		index++
	}
	if index == len(nums)-1 {
		return true
	} else {
		if nums[index] < nums[index+1] {
			asc = true
		} else {
			asc = false
		}
	}
	return sort.SliceIsSorted(nums, func(indexA int, indexB int) bool {
		if asc {
			return nums[indexA] < nums[indexB]
		} else {
			return nums[indexA] > nums[indexB]
		}
	})
}
