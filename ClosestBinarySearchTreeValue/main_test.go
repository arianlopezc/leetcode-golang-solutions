package ClosestBinarySearchTreeValue

import "testing"

func Test_closest(t *testing.T) {
	type args struct {
		root     *TreeNode
		target   *float64
		smallest *int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			closest(tt.args.root, tt.args.target, tt.args.smallest)
		})
	}
}

func Test_closestValue(t *testing.T) {
	type args struct {
		root   *TreeNode
		target float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "standard", want: 4, args: args{target: 3.714286, root: &TreeNode{Val: 4, Right: &TreeNode{Val: 5}, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}}},
		}, {
			name: "one", want: 1, args: args{target: 4.428571, root: &TreeNode{Val: 1}},
		}, {
			name: "one", want: 2, args: args{target: 3.428571, root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := closestValue(tt.args.root, tt.args.target); got != tt.want {
				t.Errorf("closestValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
