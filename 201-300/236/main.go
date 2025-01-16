package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val == p.Val || root.Val == q.Val {
		return root
	}
	val, _ := recursion(root, p, q)
	return val
}

func recursion(root, p, q *TreeNode) (out *TreeNode, findCount int) {
	if root == nil {
		return nil, 0
	}
	if root.Val == p.Val || root.Val == q.Val {
		out = root
		findCount = 1
	}

	lVal, lCount := recursion(root.Left, p, q)
	if lCount == 2 {
		return lVal, 2
	}
	rVal, rCount := recursion(root.Right, p, q)
	if rCount == 2 {
		return rVal, 2
	}
	if lCount+rCount == 2 {
		return root, 2
	}
	if findCount+lCount == 2 || findCount+rCount == 2 {
		return root, 2
	}
	if findCount == 1 || lCount == 1 || rCount == 1 {
		return nil, 1
	}
	return nil, 0
}
