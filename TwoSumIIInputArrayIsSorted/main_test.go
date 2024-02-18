package TwoSumIIInputArrayIsSorted

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	type args struct {
		numbers []int
		target  int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test", want: []int{1, 2}, args: args{target: 9, numbers: []int{2, 7, 11, 15}}},
		{name: "test", want: []int{1, 3}, args: args{target: 6, numbers: []int{2, 3, 4}}},
		{name: "test", want: []int{1, 2}, args: args{target: -1, numbers: []int{-1, 0}}},
		{name: "test", want: []int{2, 3}, args: args{target: -3, numbers: []int{-3, -2, -1}}},
		{name: "test", want: []int{2, 3}, args: args{target: 100, numbers: []int{5, 25, 75}}},
		{name: "test", want: []int{3, 6}, args: args{target: 200, numbers: []int{3, 24, 50, 79, 88, 150, 345}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.numbers, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
