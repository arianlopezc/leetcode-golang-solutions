package MaximumCountOfPositiveIntegerAndNegativeInteger

import "testing"

func Test_maximumCount(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "standard", want: 3, args: args{nums: []int{-2, -1, -1, 1, 2, 3}},
		}, {
			name: "zeros", want: 3, args: args{nums: []int{-3, -2, -1, 0, 0, 1, 2}},
		}, {
			name: "all-positives", want: 4, args: args{nums: []int{5, 20, 66, 1314}},
		}, {
			name: "all-negatives", want: 4, args: args{nums: []int{-1314, -66, -20, -5}},
		}, {
			name: "zeros-front", want: 5, args: args{nums: []int{0, 0, 1, 2, 3, 4, 5}},
		}, {
			name: "zeros-last", want: 4, args: args{nums: []int{-1314, -66, -20, -5, 0, 0}},
		}, {
			name: "zeros", want: 0, args: args{nums: []int{0, 0, 0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximumCount(tt.args.nums); got != tt.want {
				t.Errorf("maximumCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
