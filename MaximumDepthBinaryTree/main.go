package MaximumDepthBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	var queue []*TreeNode = make([]*TreeNode, 0)
	if root != nil {
		queue = append(queue, root)
	}
	var childrenQueue []*TreeNode = make([]*TreeNode, 0)
	var depth int
	for len(queue) > 0 {
		depth++
		for _, node := range queue {
			if node.Right != nil {
				childrenQueue = append(childrenQueue, node.Right)
			}
			if node.Left != nil {
				childrenQueue = append(childrenQueue, node.Left)
			}
		}
		queue = childrenQueue
		childrenQueue = make([]*TreeNode, 0)
	}
	return depth
}
