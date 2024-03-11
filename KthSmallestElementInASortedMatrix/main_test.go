package KthSmallestElementInASortedMatrix

import "testing"

func Test_kthSmallest(t *testing.T) {
	type args struct {
		matrix [][]int
		k      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "standard", want: 13, args: args{k: 8, matrix: [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}}},
		{name: "standard", want: 1, args: args{k: 1, matrix: [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}}},
		{name: "standard", want: 5, args: args{k: 2, matrix: [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}}},
		{name: "standard", want: 9, args: args{k: 3, matrix: [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}}},
		{name: "standard", want: 10, args: args{k: 4, matrix: [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}}},
		{name: "standard", want: 11, args: args{k: 5, matrix: [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}}},
		{name: "standard", want: 12, args: args{k: 6, matrix: [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}}},
		{name: "standard", want: 13, args: args{k: 7, matrix: [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}}},
		{name: "standard", want: 15, args: args{k: 9, matrix: [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}}},
		{name: "standard", want: -5, args: args{k: 1, matrix: [][]int{{-5}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.matrix, tt.args.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
