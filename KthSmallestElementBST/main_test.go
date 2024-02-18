package KthSmallestElementBST

import "testing"

func Test_findKthSmallest(t *testing.T) {
	type args struct {
		root *TreeNode
		k    int
		val  *int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "test", want: 2, args: args{k: 1, root: &TreeNode{Val: 4, Right: &TreeNode{Val: 5}, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}}},
		{name: "test", want: 3, args: args{k: 2, root: &TreeNode{Val: 4, Right: &TreeNode{Val: 5}, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}}},
		{name: "test", want: 4, args: args{k: 3, root: &TreeNode{Val: 4, Right: &TreeNode{Val: 5}, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}}},
		{name: "test", want: 5, args: args{k: 4, root: &TreeNode{Val: 4, Right: &TreeNode{Val: 5}, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 3}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.args.root, tt.args.k); got != tt.want {
				t.Errorf("findKthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
