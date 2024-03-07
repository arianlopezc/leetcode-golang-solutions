package LongestHarmoniousSubsequence

func findLHS(nums []int) int {
	var mapped = make(map[int]int)
	for _, num := range nums {
		mapped[num]++
	}
	var largestSet int
	for key, val := range mapped {
		if nextVal, ok := mapped[key+1]; ok {
			var set = val + nextVal
			if largestSet < val+nextVal {
				largestSet = set
			}
		}
	}
	return largestSet
}
