package MaximumLevelSumBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	if root == nil {
		return 1
	}
	var level = []*TreeNode{root}
	var maxValue = root.Val
	var maxLevel = 1
	var indexLevel = 1
	for len(level) > 0 {
		indexLevel++
		var children []*TreeNode
		for _, node := range level {
			if node.Left != nil {
				children = append(children, node.Left)
			}
			if node.Right != nil {
				children = append(children, node.Right)
			}
		}
		if len(children) > 0 {
			var total int
			for _, node := range children {
				total += node.Val
			}
			if total > maxValue {
				maxValue = total
				maxLevel = indexLevel
			}
		}
		level = children
	}
	return maxLevel
}
