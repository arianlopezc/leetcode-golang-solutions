package MajorityElement

import "sort"

func majorityElement(nums []int) int {
	sort.Ints(nums)
	var currentTotal, largestTotal, currentVal, largestVal = 0, 0, nums[0], nums[0]
	for _, val := range nums {
		if val == currentVal {
			currentTotal++
		} else {
			if largestTotal < currentTotal {
				largestTotal = currentTotal
				largestVal = currentVal
			}
			currentVal = val
			currentTotal = 1
		}
	}
	if largestTotal < currentTotal {
		largestVal = currentVal
	}
	return largestVal
}
