package MonotonicArray

import "testing"

func Test_isMonotonic(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "standard", want: true, args: args{nums: []int{1, 2, 2, 3}},
		}, {
			name: "standard", want: true, args: args{nums: []int{6, 5, 4, 4}},
		}, {
			name: "standard", want: true, args: args{nums: []int{1, 1, 1}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMonotonic(tt.args.nums); got != tt.want {
				t.Errorf("isMonotonic() = %v, want %v", got, tt.want)
			}
		})
	}
}
