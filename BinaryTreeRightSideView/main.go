package BinaryTreeRightSideView

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	var view []int
	if root == nil {
		return view
	}
	var level = []*TreeNode{root}
	for len(level) > 0 {
		var children []*TreeNode
		for _, node := range level {
			if node.Left != nil {
				children = append(children, node.Left)
			}
			if node.Right != nil {
				children = append(children, node.Right)
			}
		}
		view = append(view, level[len(level)-1].Val)
		level = children
	}
	return view
}
