package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	maxVal := 0
	for i := range root.Children {
		md := maxDepth(root.Children[i])
		if md+1 > maxVal {
			maxVal = md + 1
		}
	}
	if maxVal == 0 {
		return 1
	}
	return maxVal
}
