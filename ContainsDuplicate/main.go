package ContainsDuplicate

import "slices"

func containsDuplicate(nums []int) bool {
	slices.Sort(nums)
	var visited = make(map[int]bool)
	for _, value := range nums {
		if _, ok := visited[value]; ok {
			return true
		} else {
			visited[value] = true
		}
	}
	return false
}
