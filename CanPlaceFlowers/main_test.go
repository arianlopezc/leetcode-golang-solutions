package CanPlaceFlowers

import "testing"

func Test_canPlaceFlowers(t *testing.T) {
	type args struct {
		flowerbed []int
		n         int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "standard", want: true, args: args{flowerbed: []int{1, 0, 0, 0, 1}, n: 1}},
		{name: "standard fail", want: false, args: args{flowerbed: []int{1, 0, 0, 0, 1}, n: 2}},
		{name: "long", want: true, args: args{flowerbed: []int{1, 0, 0, 0, 0, 0, 0, 1}, n: 2}},
		{name: "long fail", want: false, args: args{flowerbed: []int{1, 0, 0, 0, 0, 0, 0, 1}, n: 3}},
		{name: "long impair", want: true, args: args{flowerbed: []int{0, 1, 0, 0, 0, 1, 0, 0}, n: 2}},
		{name: "long impair", want: true, args: args{flowerbed: []int{0, 0, 1, 0, 0, 1, 0, 0}, n: 2}},
		{name: "long fail", want: false, args: args{flowerbed: []int{0, 1, 0, 0, 0, 1, 0, 0}, n: 3}},
		{name: "short", want: true, args: args{flowerbed: []int{0, 0, 1, 0, 0}, n: 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPlaceFlowers(tt.args.flowerbed, tt.args.n); got != tt.want {
				t.Errorf("canPlaceFlowers() = %v, want %v", got, tt.want)
			}
		})
	}
}
