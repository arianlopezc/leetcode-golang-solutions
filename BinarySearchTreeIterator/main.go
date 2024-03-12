package BinarySearchTreeIterator

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func transform(root *TreeNode, tail *TreeNode) *TreeNode {
	if root == nil {
		return tail
	}
	res := transform(root.Left, root)
	root.Left = nil
	root.Right = transform(root.Right, tail)
	return res
}

type BSTIterator struct {
	Node *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	var start = transform(root, nil)
	return BSTIterator{Node: &TreeNode{Right: start}}
}

func (this *BSTIterator) Next() int {
	this.Node = this.Node.Right
	var res = this.Node.Val
	return res
}

func (this *BSTIterator) HasNext() bool {
	if this.Node != nil && this.Node.Right != nil {
		return true
	}
	return false
}
