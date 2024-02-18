package BinaryTreeVerticalOrderTRaversal

import (
	"reflect"
	"testing"
)

func Test_verticalOrder(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{name: "standard", want: [][]int{{4}, {9, 5}, {3, 0, 1}, {8, 2}, {7}}, args: args{root: &TreeNode{Val: 3, Left: &TreeNode{Val: 9, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 0, Right: &TreeNode{Val: 2}}}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 7}}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := verticalOrder(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("verticalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
