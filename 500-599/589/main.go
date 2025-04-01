package main

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	if root == nil {
		return nil
	}
	out := make([]int, 0)
	out = append(out, root.Val)
	for i := range root.Children {
		out = append(out, preorder(root.Children[i])...)
	}
	return out
}
