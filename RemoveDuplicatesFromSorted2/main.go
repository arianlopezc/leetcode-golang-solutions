package RemoveDuplicatesFromSorted2

func removeDuplicates(nums []int) int {
	var length = len(nums)
	if length <= 1 {
		return length
	}
	var inDuplex = false
	var index, writeIndex = 0, 0
	for index < length {
		if index+1 < length && nums[index] == nums[index+1] {
			if !inDuplex {
				inDuplex = true
				nums[writeIndex] = nums[index]
				writeIndex++
			}
		} else {
			inDuplex = false
			nums[writeIndex] = nums[index]
			writeIndex++
		}
		index++
	}
	return writeIndex
}
