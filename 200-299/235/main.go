package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return recursion(root, p, q)
}

func recursion(root, p, q *TreeNode) *TreeNode {
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	a := root.Val < p.Val
	b := root.Val < q.Val
	if a && b {
		return recursion(root.Right, p, q)
	}
	if !a && !b {
		return recursion(root.Left, p, q)
	}
	return root
}
