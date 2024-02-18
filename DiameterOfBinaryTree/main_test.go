package DiameterOfBinaryTree

import "testing"

func Test_diameter(t *testing.T) {
	type args struct {
		largest *int
		root    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diameter(tt.args.largest, tt.args.root); got != tt.want {
				t.Errorf("diameter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_diameterOfBinaryTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one", want: 1, args: args{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}}},
		}, {
			name: "three", want: 3, args: args{root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := diameterOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("diameterOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
