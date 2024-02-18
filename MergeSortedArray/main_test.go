package MergeSortedArray

import (
	"testing"
)

func Test_main(t *testing.T) {
	tests := []struct {
		name         string
		nums1, nums2 []int
		m, n         int
	}{
		{name: "regular", m: 3, n: 3, nums1: []int{1, 2, 3, 0, 0, 0}, nums2: []int{2, 5, 6}},
		{name: "m1", m: 1, n: 0, nums1: []int{1}, nums2: []int{}},
		{name: "m2", m: 0, n: 1, nums1: []int{0}, nums2: []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.nums1, tt.m, tt.nums2, tt.n)
		})
	}
}
