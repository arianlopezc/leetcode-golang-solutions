package TwoSumIVInputIsaBST

import "testing"

func Test_findTarget(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "standard", want: true, args: args{k: 9, root: &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}}},
		},
		{
			name: "standard", want: true, args: args{k: 10, root: &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}}},
		},
		{
			name: "standard", want: false, args: args{k: 28, root: &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}}},
		},
		{
			name: "standard", want: false, args: args{k: 1, root: &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 6, Right: &TreeNode{Val: 7}}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findTarget(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("findTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
