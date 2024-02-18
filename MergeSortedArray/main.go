package MergeSortedArray

func merge(nums1 []int, m int, nums2 []int, n int) {
	for index := len(nums1) - 1; index >= 0; index-- {
		if m > 0 && n > 0 {
			if nums1[m-1] <= nums2[n-1] {
				nums1[index] = nums2[n-1]
				n--
			} else {
				nums1[index] = nums1[m-1]
				m--
			}
		} else if m == 0 {
			nums1[index] = nums2[n-1]
			n--
		}
	}
}
