package MinimumSizeSubarraySum

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	type args struct {
		target int
		nums   []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", want: 2, args: args{nums: []int{2, 3, 1, 2, 4, 3}, target: 7}},
		{name: "test", want: 1, args: args{nums: []int{4, 1, 4}, target: 4}},
		{name: "test", want: 0, args: args{nums: []int{1, 1, 1, 1, 1, 1, 1, 1}, target: 11}},
		{name: "test", want: 3, args: args{nums: []int{1, 2, 3, 4, 5}, target: 11}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.args.target, tt.args.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
