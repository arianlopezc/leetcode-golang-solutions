package LowestCommonAncestor

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func lowestCommonAncestor(p *Node, q *Node) *Node {
	var visited = make(map[int]bool)
	for p != nil {
		visited[p.Val] = true
		p = p.Parent
	}
	for q != nil && !visited[q.Val] {
		visited[q.Val] = true
		q = q.Parent
	}
	return q.Parent
}
