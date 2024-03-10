package MissingRanges

func findMissingRanges(nums []int, lower int, upper int) [][]int {
	var length = len(nums)
	var result [][]int
	for index := 0; index < length; index++ {
		if lower != nums[index] {
			result = append(result, []int{lower, nums[index] - 1})
		}
		lower = nums[index] + 1
	}
	if lower <= upper {
		result = append(result, []int{lower, upper})
	}
	return result
}
