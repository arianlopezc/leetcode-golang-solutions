package FibonacciNumber

import "testing"

func Test_fib(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "two", want: 1, args: args{n: 2},
		}, {
			name: "zero", want: 0, args: args{n: 0},
		}, {
			name: "three", want: 2, args: args{n: 3},
		}, {
			name: "fourth", want: 3, args: args{n: 4},
		}, {
			name: "fourth", want: 5, args: args{n: 5},
		}, {
			name: "fourth", want: 8, args: args{n: 6},
		}, {
			name: "fourth", want: 13, args: args{n: 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib(tt.args.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
