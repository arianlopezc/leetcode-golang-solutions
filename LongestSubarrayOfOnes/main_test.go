package LongestSubarrayOfOnes

import "testing"

func Test_longestSubarray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", want: 3, args: args{nums: []int{1, 1, 0, 1}}},
		{name: "test", want: 5, args: args{nums: []int{0, 1, 1, 1, 0, 1, 1, 0, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestSubarray(tt.args.nums); got != tt.want {
				t.Errorf("longestSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
