package FindDifferences

import (
	"reflect"
	"testing"
)

func Test_findDifference(t *testing.T) {
	type args struct {
		nums1 []int
		nums2 []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "standard", want: [][]int{[]int{1, 3}, []int{4, 6}}, args: args{nums1: []int{1, 2, 3}, nums2: []int{2, 4, 6}}},
		{name: "standard", want: [][]int{[]int{33, -20, 5, 26, 29}, []int{0, -34, -67, -36, -84, -100, -69, 52}}, args: args{nums1: []int{80, 5, -20, 33, 26, 29}, nums2: []int{-69, 0, -36, 52, -84, -34, -67, -100, 80}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findDifference(tt.args.nums1, tt.args.nums2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
