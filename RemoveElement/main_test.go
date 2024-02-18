package RemoveElement

import "testing"

func Test_removeElement(t *testing.T) {
	tests := []struct {
		name string
		want int
		nums []int
		val  int
	}{
		{name: "regular", nums: []int{3, 2, 2, 3}, val: 2, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeElement(tt.nums, tt.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
