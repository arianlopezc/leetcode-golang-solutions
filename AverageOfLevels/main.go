package AverageOfLevels

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	var queue = make([]*TreeNode, 0)
	queue = append(queue, root)
	var childrenQueue = make([]*TreeNode, 0)
	var result = make([]float64, 0)
	for len(queue) > 0 {
		var ave int
		var total int
		for _, node := range queue {
			total++
			ave += node.Val
			if node.Right != nil {
				childrenQueue = append(childrenQueue, node.Right)
			}
			if node.Left != nil {
				childrenQueue = append(childrenQueue, node.Left)
			}
		}
		result = append(result, float64(total)/float64(ave))
		queue = childrenQueue
		childrenQueue = make([]*TreeNode, 0)
	}
	return result
}
