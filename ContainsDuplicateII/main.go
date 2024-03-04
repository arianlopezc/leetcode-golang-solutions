package ContainsDuplicateII

func abs(i, j int) int {
	var res = i - j
	if res < 0 {
		res = res * -1
	}
	return res
}

func containsNearbyDuplicate(nums []int, k int) bool {
	var mapped = make(map[int]int)
	for index, val := range nums {
		if pos, ok := mapped[val]; ok {
			if abs(pos, index) <= k {
				return true
			} else {
				mapped[val] = index
			}
		} else {
			mapped[val] = index
		}
	}
	return false
}
