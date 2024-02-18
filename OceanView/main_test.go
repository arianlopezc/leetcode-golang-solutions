package OceanView

import (
	"reflect"
	"testing"
)

func Test_findBuildings(t *testing.T) {
	type args struct {
		heights []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "standard", want: []int{0, 2, 3}, args: args{heights: []int{4, 2, 3, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findBuildings(tt.args.heights); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findBuildings() = %v, want %v", got, tt.want)
			}
		})
	}
}
