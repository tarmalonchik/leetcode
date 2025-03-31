package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	a, b := diameterRecursion(root)
	return getMax(a, b) - 1
}

func diameterRecursion(root *TreeNode) (extendable, nonExtendable int) {
	if root.Left == nil && root.Right == nil {
		return 1, 0
	}
	if root.Left != nil && root.Right == nil {
		ext, nonExt := diameterRecursion(root.Left)
		return ext + 1, nonExt
	}
	if root.Right != nil && root.Left == nil {
		ext, nonExt := diameterRecursion(root.Right)
		return ext + 1, nonExt
	}

	lExt, lNonExt := diameterRecursion(root.Left)
	rExt, rNonExt := diameterRecursion(root.Right)
	nonExtendable = getMax3(lExt+rExt+1, lNonExt, rNonExt)
	extendable = getMax(lExt, rExt) + 1
	return extendable, nonExtendable
}

func getMax3(a, b, c int) int {
	return getMax(getMax(a, b), c)
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
