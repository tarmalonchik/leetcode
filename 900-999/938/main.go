package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	out := 0
	out += rangeSumBST(root.Left, low, high)
	out += rangeSumBST(root.Right, low, high)
	if root.Val >= low && root.Val <= high {
		out += root.Val
	}
	return out
}
