package ConvertBinarySearchTreeToSortedDoublyLinkedList

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func flatten(node *Node, tail *Node) *Node {
	if node == nil {
		return tail
	}
	var res = flatten(node.Left, node)
	node.Left = nil
	node.Right = flatten(node.Right, tail)
	return res
}

func connect(node *Node) {
	var nav = node
	for nav.Right != nil {
		nav.Right.Left = nav
		nav = nav.Right
	}
	node.Left = nav
	nav.Right = node
}

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return root
	}
	var smallestNode = flatten(root, nil)
	connect(smallestNode)
	return smallestNode
}
