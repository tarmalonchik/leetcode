package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeaves(root *TreeNode) int {
	return recursion(root, false)
}

func recursion(node *TreeNode, isLeft bool) int {
	if node.Left == nil && node.Right == nil {
		if isLeft {
			return node.Val
		}
		return 0
	}
	out := 0
	if node.Left != nil {
		out += recursion(node.Left, true)
	}
	if node.Right != nil {
		out += recursion(node.Right, false)
	}
	return out
}
