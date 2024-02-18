package SuccessfulPairsSpellsPotions

import (
	"reflect"
	"testing"
)

func Test_successfulPairs(t *testing.T) {
	type args struct {
		spells  []int
		potions []int
		success int64
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test", want: []int{4, 0, 3}, args: args{success: 7, spells: []int{5, 1, 3}, potions: []int{1, 2, 3, 4, 5}}},
		{name: "test", want: []int{1, 0, 0}, args: args{success: 25, spells: []int{5, 1, 3}, potions: []int{1, 2, 3, 4, 5}}},
		{name: "test", want: []int{2, 0, 2}, args: args{success: 16, spells: []int{3, 1, 2}, potions: []int{8, 5, 8}}},
		{name: "test", want: []int{1}, args: args{success: 10, spells: []int{7}, potions: []int{2, 1}}},
		{name: "test", want: []int{2, 6}, args: args{success: 320, spells: []int{9, 39}, potions: []int{35, 40, 22, 37, 29, 22}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := successfulPairs(tt.args.spells, tt.args.potions, tt.args.success); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("successfulPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
