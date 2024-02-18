package KthMissingPositiveNumber

import "testing"

func Test_findKthPositive(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "regular", want: 12, args: args{k: 7, arr: []int{2, 3, 4, 7, 11}},
		}, {
			name: "sequence", want: 11, args: args{k: 7, arr: []int{2, 4, 6, 8}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthPositive(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("findKthPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
