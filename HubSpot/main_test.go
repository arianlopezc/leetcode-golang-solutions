package HubSpot

import (
	"reflect"
	"testing"
)

func Test_exercise(t *testing.T) {
	type args struct {
		a         []int
		b         []int
		maxLength int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "regular", args: args{a: []int{1, 3, 5}, b: []int{2, 6, 8, 9}, maxLength: 3}, want: []int{1, 2, 3}},
		{name: "duplicates", args: args{a: []int{1, 1, 3, 5}, b: []int{2, 2, 8, 9}, maxLength: 5}, want: []int{1, 1, 2, 2, 3}},
		{name: "smallest", args: args{a: []int{1, 3, 5}, b: []int{2, 6, 8, 9}, maxLength: 5}, want: []int{1, 2, 3, 5, 6}},
		{name: "emptyA", args: args{a: []int{}, b: []int{2, 6, 8, 9}, maxLength: 4}, want: []int{2, 6, 8, 9}},
		{name: "emptyB", args: args{a: []int{1, 3, 5}, b: []int{}, maxLength: 3}, want: []int{1, 3, 5}},
		{name: "over limit length", args: args{a: []int{1, 3, 5}, b: []int{2, 6, 8, 9}, maxLength: 100}, want: []int{1, 2, 3, 5, 6, 8, 9}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := exercise(tt.args.a, tt.args.b, tt.args.maxLength); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("exercise() = %v, want %v", got, tt.want)
			}
		})
	}
}
