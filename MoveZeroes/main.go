package MoveZeroes

func moveZeroes(nums []int) {
	var length int = len(nums)
	if length > 1 {
		var zeroIndex int = 0
		var nonZeroIndex int = 1
		var zeroIndexStopAt int = length - 2
		for zeroIndex <= zeroIndexStopAt {
			if nums[zeroIndex] != 0 {
				zeroIndex++
				nonZeroIndex = zeroIndex + 1
			} else if nonZeroIndex < length {
				if nums[nonZeroIndex] == 0 {
					nonZeroIndex++
				} else {
					nums[zeroIndex], nums[nonZeroIndex] = nums[nonZeroIndex], nums[zeroIndex]
					zeroIndex++
				}
			} else {
				break
			}
		}
	}
}
