package TopKFrequentElements

import (
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	var instances = make(map[int]int)
	for _, val := range nums {
		instances[val]++
	}
	var keys []int
	for key := range instances {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(a int, b int) bool {
		if instances[keys[a]] < instances[keys[b]] {
			return false
		} else {
			return true
		}
	})
	return keys[:k]
}
