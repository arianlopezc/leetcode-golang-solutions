package FindPeakElement

func findPeakElement(nums []int) int {
	var length = len(nums)
	if length == 1 {
		return 0
	}
	var l, r = 0, length
	for l < r {
		middle := (l + r) / 2
		if left, right := middle-1 >= 0, middle+1 < length; (!left && !right) || (!left && right && nums[middle] > nums[middle+1]) || (left && !right && nums[middle] > nums[middle-1]) || (left && right && nums[middle] > nums[middle+1] && nums[middle] > nums[middle-1]) {
			return middle
		} else if left && nums[middle-1] > nums[middle] {
			r = middle - 1
		} else if right && nums[middle+1] > nums[middle] {
			l = middle + 1
		}
	}
	return l
}
