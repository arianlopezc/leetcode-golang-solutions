package SubarraySumEqualsK

func subarraySum(nums []int, k int) int {
	var count, sum int
	prefixSums := map[int]int{0: 1}
	for _, n := range nums {
		sum += n
		count += prefixSums[sum-k]
		prefixSums[sum]++
	}
	return count
}
