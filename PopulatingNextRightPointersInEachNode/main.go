package PopulatingNextRightPointersInEachNode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root != nil {
		var children = []*Node{root}
		for len(children) > 0 {
			var temp []*Node
			for index := 0; index < len(children); index++ {
				if children[index].Left != nil {
					temp = append(temp, children[index].Left)
				}
				if children[index].Right != nil {
					temp = append(temp, children[index].Right)
				}
				if index < len(children)-1 {
					children[index].Next = children[index+1]
				}
			}
			children = temp
		}
	}
	return root
}
