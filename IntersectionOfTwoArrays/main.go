package IntersectionOfTwoArrays

func intersection(nums1 []int, nums2 []int) []int {
	var intersected []int
	var has = make(map[int]int)
	for _, val := range nums1 {
		if _, ok := has[val]; !ok {
			has[val] = 1
		}
	}
	for _, val := range nums2 {
		if tot, ok := has[val]; ok && tot == 1 {
			intersected = append(intersected, val)
			has[val]++
		}
	}
	return intersected
}
