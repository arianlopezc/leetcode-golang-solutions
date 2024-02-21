package InOrderTraversalStack

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "standard", want: []int{1, 3, 2}, args: args{root: &TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}}},
		{name: "one", want: []int{1}, args: args{root: &TreeNode{Val: 1}}},
		{name: "none", want: []int{}, args: args{root: nil}},
		{name: "standard", want: []int{4, 9, 0, 2, 3, 5, 1, 8, 7}, args: args{root: &TreeNode{Val: 3,
			Left: &TreeNode{Val: 9,
				Left: &TreeNode{Val: 4},
				Right: &TreeNode{Val: 0,
					Right: &TreeNode{Val: 2},
				},
			},
			Right: &TreeNode{Val: 8,
				Left: &TreeNode{Val: 1,
					Left: &TreeNode{Val: 5},
				},
				Right: &TreeNode{Val: 7},
			},
		},
		},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
