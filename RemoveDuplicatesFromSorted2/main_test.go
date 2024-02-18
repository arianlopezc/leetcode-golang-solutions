package RemoveDuplicatesFromSorted2

import "testing"

func Test_removeDuplicates(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "duplicated long", want: 8, args: args{nums: []int{0, 0, 1, 1, 1, 1, 2, 2, 3, 3}}},
		{name: "duplicated small", want: 5, args: args{nums: []int{1, 1, 1, 2, 2, 3}}},
		{name: "regular", want: 9, args: args{nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}}},
		{name: "ordered", want: 6, args: args{nums: []int{0, 1, 2, 3, 4, 5}}},
		{name: "one", want: 1, args: args{nums: []int{0}}},
		{name: "two-double", want: 2, args: args{nums: []int{1, 1}}},
		{name: "two-double", want: 2, args: args{nums: []int{1, 1, 1}}},
		{name: "two-double", want: 5, args: args{nums: []int{1, 1, 1, 1, 2, 2, 3}}},
		{name: "some", want: 5, args: args{nums: []int{1, 1, 1, 1, 2, 2, 3}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicates(tt.args.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
