package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	l := searchBST(root.Left, val)
	if l != nil {
		return l
	}

	r := searchBST(root.Right, val)
	return r
}
