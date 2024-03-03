package IncreasingOrderSearchTree

import (
	"reflect"
	"testing"
)

func Test_increasingBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "test", args: args{root: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}}, want: &TreeNode{Val: 1, Right: &TreeNode{Val: 1, Right: &TreeNode{Val: 3}}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := increasingBST(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("increasingBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
