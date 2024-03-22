package NonDecreasingArray

func checkPossibility(nums []int) bool {
	var invalidIndex = -1
	var index = 0
	for index < len(nums)-2 {
		if nums[index] > nums[index+1] {
			if invalidIndex == -1 {
				invalidIndex = index
			} else {
				return false
			}
		}
		index++
	}
	if invalidIndex <= 0 || invalidIndex == len(nums)-2 || nums[invalidIndex] <= nums[invalidIndex+2] || nums[invalidIndex-1] <= nums[invalidIndex+1] {
		return true
	}
	return false
}
