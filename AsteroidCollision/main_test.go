package AsteroidCollision

import (
	"reflect"
	"testing"
)

func Test_asteroidCollision(t *testing.T) {
	type args struct {
		asteroids []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "test", want: []int{5, 10}, args: args{asteroids: []int{5, 10, -5}}},
		{name: "test", want: []int{}, args: args{asteroids: []int{8, -8}}},
		{name: "test", want: []int{10}, args: args{asteroids: []int{10, 2, -5}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := asteroidCollision(tt.args.asteroids); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("asteroidCollision() = %v, want %v", got, tt.want)
			}
		})
	}
}
