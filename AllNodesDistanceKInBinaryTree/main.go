package AllNodesDistanceKInBinaryTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func down(node *TreeNode, res *[]int, level int, goal int) {
	if node != nil {
		if level == goal {
			*res = append(*res, node.Val)
		} else {
			if level < goal {
				down(node.Left, res, level+1, goal)
				down(node.Right, res, level+1, goal)
			}
		}
	}
}

func navigate(node *TreeNode, target int, res *[]int, goal int) int {
	if node != nil {
		if node.Val == target {
			down(node, res, 0, goal)
			return 1
		} else {
			var leftNav = navigate(node.Left, target, res, goal)
			if leftNav != -1 {
				if leftNav == goal {
					*res = append(*res, node.Val)
				} else {
					down(node.Right, res, leftNav+1, goal)
				}
				return leftNav + 1
			} else {
				var rightNav = navigate(node.Right, target, res, goal)
				if rightNav != -1 {
					if rightNav == goal {
						*res = append(*res, node.Val)
					} else {
						down(node.Left, res, rightNav+1, goal)
					}
					return rightNav + 1
				}
			}
			return -1
		}
	} else {
		return -1
	}
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	var res []int = make([]int, 0)
	navigate(root, target.Val, &res, k)
	return res
}
