package BinaryTreeLongestConsecutiveSequence

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func consecutive(node *TreeNode, currMax int, maxPath *int) {
	if node != nil {
		if node.Left != nil && node.Left.Val == node.Val+1 {
			if currMax+1 > *maxPath {
				*maxPath = currMax + 1
			}
			consecutive(node.Left, currMax+1, maxPath)
		} else {
			consecutive(node.Left, 1, maxPath)
		}
		if node.Right != nil && node.Right.Val == node.Val+1 {
			if currMax+1 > *maxPath {
				*maxPath = currMax + 1
			}
			consecutive(node.Right, currMax+1, maxPath)
		} else {
			consecutive(node.Right, 1, maxPath)
		}
	}
}

func longestConsecutive(root *TreeNode) int {
	var maxPath = 1
	consecutive(root, 1, &maxPath)
	return maxPath
}
