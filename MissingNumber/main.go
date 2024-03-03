package MissingNumber

func missingNumber(nums []int) int {
	n := len(nums)
	expectedSum := n * (n + 1) / 2
	var actualSum int
	for _, val := range nums {
		actualSum += val
	}
	return expectedSum - actualSum
}
