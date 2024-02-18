package MinimumAbsoluteDifferenceBST

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func absComp(val int) int {
	if val < 0 {
		return val * -1
	}
	return val
}

func getMinimumDifference(root *TreeNode) int {
	var abs, absLeft, absRight int
	if root.Left != nil {
		absLeft = absComp(root.Val - root.Left.Val)
	} else {
		absLeft = 1000000
	}
	if root.Right != nil {
		absRight = absComp(root.Val - root.Right.Val)
	} else {
		absRight = 1000000
	}
	if absLeft < absRight {
		abs = absLeft
	} else {
		abs = absRight
	}
	var absChildLeft, absChildRight int
	if root.Left != nil && (root.Left.Left != nil || root.Left.Right != nil) {
		absChildLeft = getMinimumDifference(root.Left)
		if abs > absChildLeft {
			abs = absChildLeft
		}
	}
	if root.Right != nil && (root.Right.Left != nil || root.Right.Right != nil) {
		absChildRight = getMinimumDifference(root.Right)
		if abs > absChildRight {
			abs = absChildRight
		}
	}
	return abs
}
