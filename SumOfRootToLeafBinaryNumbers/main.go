package SumOfRootToLeafBinaryNumbers

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	var total int
	if root != nil {
		sum(root, &total, "")
	}
	return total
}

func sum(node *TreeNode, total *int, path string) {
	if node.Left == nil && node.Right == nil {
		i, _ := strconv.ParseInt(path+strconv.Itoa(node.Val), 2, 64)
		*total = *total + int(i)
	} else {
		if node.Left != nil {
			sum(node.Left, total, path+strconv.Itoa(node.Val))
		}
		if node.Right != nil {
			sum(node.Right, total, path+strconv.Itoa(node.Val))
		}
	}
}
