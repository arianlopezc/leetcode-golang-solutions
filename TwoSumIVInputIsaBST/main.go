package TwoSumIVInputIsaBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	var sortedElements = make([]int, 0)
	sort(root, &sortedElements)
	var left, right = 0, len(sortedElements) - 1
	for left < right {
		var opResult = sortedElements[left] + sortedElements[right]
		if k == opResult {
			return true
		} else if k > opResult {
			left++
		} else if k < opResult {
			right--
		}
	}
	return false
}

func sort(node *TreeNode, sortedElements *[]int) {
	if node != nil {
		sort(node.Left, sortedElements)
		*sortedElements = append(*sortedElements, node.Val)
		sort(node.Right, sortedElements)
	}
}
