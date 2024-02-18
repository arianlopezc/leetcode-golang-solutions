package BinaryTreeVerticalOrderTRaversal

//import "slices"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func verticalOrder(root *TreeNode) [][]int {
	var mapped = make(map[int][]int)
	var result [][]int
	if root == nil {
		return result
	}
	var pendingNode = []*TreeNode{root}
	var pendingLevel = []int{0}
	var smallest, highest = 0, 0
	for len(pendingNode) > 0 {
		var newPendingNodes []*TreeNode
		var newPendingLevels []int
		for index, node := range pendingNode {
			mapped[pendingLevel[index]] = append(mapped[pendingLevel[index]], node.Val)
			if node.Left != nil {
				newPendingNodes = append(newPendingNodes, node.Left)
				newPendingLevels = append(newPendingLevels, pendingLevel[index]-1)
				smallest = min(smallest, pendingLevel[index]-1)
			}
			if node.Right != nil {
				newPendingNodes = append(newPendingNodes, node.Right)
				newPendingLevels = append(newPendingLevels, pendingLevel[index]+1)
				highest = max(highest, pendingLevel[index]+1)
			}
		}
		pendingNode = newPendingNodes
		pendingLevel = newPendingLevels
	}
	for smallest <= highest {
		var list = mapped[smallest]
		if len(list) > 0 {
			result = append(result, list)
		} else {
			return result
		}
		smallest++
	}
	return result
}
