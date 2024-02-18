package FindPeakElement

import "testing"

func Test_findPeakElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", want: 0, args: args{nums: []int{2, 1}}},
		{name: "test", want: 1, args: args{nums: []int{1, 2}}},
		{name: "test", want: 2, args: args{nums: []int{1, 2, 3, 1}}},
		{name: "test", want: 5, args: args{nums: []int{1, 2, 1, 3, 5, 6, 4}}},
		{name: "test", want: 5, args: args{nums: []int{1, 2, 1, 3, 5, 6, 4}}},
		{name: "test", want: 10, args: args{nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}}},
		{name: "test", want: 0, args: args{nums: []int{-1, -2, -3, -4, -5, -6, -7, -8, -9, -10, -11}}},
		{name: "test", want: 0, args: args{nums: []int{11, 10, 9, 8, 7, 6, 5, 4, 3, 2, 1}}},
		{name: "test", want: 10, args: args{nums: []int{-11, -10, -9, -8, -7, -6, -5, -4, -3, -2, -1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findPeakElement(tt.args.nums); got != tt.want {
				t.Errorf("findPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
