package MaxNumOfKSumPairs

func maxOperations(nums []int, k int) int {
	var totalOps int
	mapped := make(map[int]int)
	for _, val := range nums {
		diff := k - val
		if diff > 0 {
			if mapped[diff] > 0 {
				totalOps++
				mapped[diff] = mapped[diff] - 1
			} else {
				mapped[val] = mapped[val] + 1
			}
		}
	}
	return totalOps
}
