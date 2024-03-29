package SumRootToLeafNumbers

import "testing"

func Test_sumNumbers(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", want: 25, args: args{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}}},
		{name: "test", want: 1026, args: args{root: &TreeNode{Val: 4, Left: &TreeNode{Val: 9, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 1}}, Right: &TreeNode{Val: 0}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumNumbers(tt.args.root); got != tt.want {
				t.Errorf("sumNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
