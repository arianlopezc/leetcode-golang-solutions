package ConstructBinaryTreeFromString

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func build(s string, index int) (i int, node *TreeNode) {
	var val []rune
	for index < len(s) {
		if s[index] != '(' && s[index] != ')' {
			val = append(val, rune(s[index]))
			index++
		} else {
			if node == nil {
				res, _ := strconv.Atoi(string(val))
				node = &TreeNode{Val: res}
			}
			if s[index] == '(' {
				if node.Left == nil {
					index, node.Left = build(s, index+1)
				} else {
					index, node.Right = build(s, index+1)
				}
			} else {
				return index + 1, node
			}
		}
	}
	if node == nil {
		res, _ := strconv.Atoi(string(val))
		node = &TreeNode{Val: res}
	}
	return index, node
}

func str2tree(s string) *TreeNode {
	if len(s) == 0 {
		return nil
	}
	_, node := build(s, 0)
	return node
}
