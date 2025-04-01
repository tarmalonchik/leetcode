package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	return isSubtreeRecursion(root, subRoot, false)
}

func isSubtreeRecursion(root *TreeNode, subRoot *TreeNode, isPart bool) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if root == nil || subRoot == nil {
		return false
	}
	if root.Val == subRoot.Val {
		if isSubtreeRecursion(root.Left, subRoot.Left, true) && isSubtreeRecursion(root.Right, subRoot.Right, true) {
			return true
		}
	}
	if isPart {
		return false
	}
	return isSubtreeRecursion(root.Left, subRoot, false) || isSubtreeRecursion(root.Right, subRoot, false)
}
