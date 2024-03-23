package MaximumWidthOfBinaryTree

import "testing"

func Test_widthOfBinaryTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "standard",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val:   5,
							Left:  nil,
							Right: nil,
						},
					},
					Right: &TreeNode{
						Val:  2,
						Left: nil,
					},
				},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := widthOfBinaryTree(tt.args.root); got != tt.want {
				t.Errorf("widthOfBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
