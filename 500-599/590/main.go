package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return nil
	}
	out := make([]int, 0)
	for i := range root.Children {
		out = append(out, postorder(root.Children[i])...)
	}
	out = append(out, root.Val)
	return out
}
