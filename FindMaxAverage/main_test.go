package FindMaxAverage

import "testing"

func Test_findMaxAverage(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{name: "regular", want: 12.75000, args: args{nums: []int{1, 12, -5, -6, 50, 3}, k: 4}},
		{name: "zero", want: 0, args: args{nums: []int{1, 12, -5, -6, 50, 3}, k: 0}},
		{name: "small", want: 5.00000, args: args{nums: []int{5}, k: 1}},
		{name: "negative", want: -1, args: args{nums: []int{-2, -1}, k: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMaxAverage(tt.args.nums, tt.args.k); got != tt.want {
				t.Errorf("findMaxAverage() = %v, want %v", got, tt.want)
			}
		})
	}
}
