package FindModeInBinarySearchTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	var nodes = []*TreeNode{root}
	var totals = make(map[int]int)
	var mode = 0
	for len(nodes) > 0 {
		var node = nodes[0]
		totals[node.Val]++
		if mode < totals[node.Val] {
			mode = totals[node.Val]
		}
		if node.Left != nil {
			nodes = append(nodes, node.Left)
		}
		if node.Right != nil {
			nodes = append(nodes, node.Right)
		}
		nodes = nodes[1:]
	}
	var result []int
	for key, total := range totals {
		if total == mode {
			result = append(result, key)
		}
	}
	return result
}
