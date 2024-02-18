package BinaryTreeLevelOrderTraversal

import "reflect"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverse(s interface{}) {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var queue = []*TreeNode{root}
	var result [][]int
	var direction = false
	for len(queue) > 0 {
		var children []*TreeNode
		var level []int
		for _, node := range queue {
			level = append(level, node.Val)
			if node.Left != nil {
				children = append(children, node.Left)
			}
			if node.Right != nil {
				children = append(children, node.Right)
			}
		}
		result = append(result, level)
		if direction {
			reverse(level)
		}
		queue = children
		direction = !direction
	}
	return result
}
