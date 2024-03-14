package AllElementsInTwoBinarySearchTrees

import "sort"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func navigate(node *TreeNode, res *[]int) {
	if node != nil {
		navigate(node.Left, res)
		*res = append(*res, node.Val)
		navigate(node.Right, res)
	}
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	var sortedTreeOne = make([]int, 0)
	var sortedTreeTwo = make([]int, 0)
	navigate(root1, &sortedTreeOne)
	navigate(root2, &sortedTreeTwo)
	sortedTreeOne = append(sortedTreeOne, sortedTreeTwo...)
	sort.Ints(sortedTreeOne)
	return sortedTreeOne
}
