package FindDifferences

func findDifference(nums1 []int, nums2 []int) [][]int {
	var mappedNumsOne = make(map[int]bool)
	var mappedNumsTwo = make(map[int]bool)
	for _, value := range nums1 {
		mappedNumsOne[value] = true
	}
	for _, value := range nums2 {
		mappedNumsTwo[value] = true
	}
	var missingInTwo []int
	for head := range mappedNumsOne {
		if !mappedNumsTwo[head] {
			missingInTwo = append(missingInTwo, head)
		}
	}
	var missingInOne []int
	for head := range mappedNumsTwo {
		if !mappedNumsOne[head] {
			missingInOne = append(missingInOne, head)
		}
	}
	return [][]int{missingInTwo, missingInOne}
}
