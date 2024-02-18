package ProductOfArrayExceptSelf

import (
	"reflect"
	"testing"
)

func Test_productExceptSelf(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "standard", want: []int{24, 12, 8, 6}, args: args{nums: []int{1, 2, 3, 4}}},
		{name: "standard", want: []int{0, 0, 9, 0, 0}, args: args{nums: []int{-1, 1, 0, -3, 3}}},
		{name: "standard", want: []int{1, -1}, args: args{nums: []int{-1, 1}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := productExceptSelf(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("productExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}
