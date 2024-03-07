package BinaryTreePaths

import (
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func navigate(node *TreeNode, path string, res *[]string) {
	if node.Left == nil && node.Right == nil {
		path += strconv.Itoa(node.Val)
		*res = append(*res, path)
	} else {
		path += strconv.Itoa(node.Val)
		if node.Left != nil {
			navigate(node.Left, path+"->", res)
		}
		if node.Right != nil {
			navigate(node.Right, path+"->", res)
		}
	}
}

func binaryTreePaths(root *TreeNode) []string {
	var res []string
	if root != nil {
		navigate(root, "", &res)
	}
	return res
}
