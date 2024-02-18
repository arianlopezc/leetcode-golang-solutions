package MaxConsecutiveOnes

func longestOnes(nums []int, k int) int {
	var totalCanAdd = k
	var onesInWindow int
	var maxAdded int
	var index int
	var length = len(nums)
	var rearIndex = 0
	for index < length {
		if totalCanAdd == 0 && nums[index] == 0 {
			if nums[rearIndex] == 0 {
				totalCanAdd++
			}
			onesInWindow--
			rearIndex++
		}
		if nums[index] == 1 || totalCanAdd > 0 {
			onesInWindow++
			if onesInWindow > maxAdded {
				maxAdded = onesInWindow
			}
			if nums[index] == 0 {
				totalCanAdd--
			}
			index++
		}
	}
	return maxAdded
}
