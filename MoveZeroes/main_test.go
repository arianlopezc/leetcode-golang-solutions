package MoveZeroes

import "testing"

func Test_moveZeroes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "standard", args: args{nums: []int{0, 1, 0, 3, 12}}},
		{name: "zero", args: args{nums: []int{0, 0, 0, 0}}},
		{name: "standard", args: args{nums: []int{0, 1, 0, 3, 0}}},
		{name: "standard", args: args{nums: []int{12, 0, 0, 0, 0}}},
		{name: "standard", args: args{nums: []int{12, 1, 1, 1, 32}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			moveZeroes(tt.args.nums)
		})
	}
}
