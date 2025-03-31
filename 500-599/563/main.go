package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	if root == nil {
		return 0
	}
	_, tilt := findTiltRecursive(root)
	return tilt
}

func findTiltRecursive(root *TreeNode) (sum, tilt int) {
	if root.Left == nil && root.Right == nil {
		return root.Val, 0
	}

	if root.Left != nil && root.Right == nil {
		sum, tilt = findTiltRecursive(root.Left)
		tilt += abs(sum, 0)
		sum += root.Val
		return sum, tilt
	}

	if root.Right != nil && root.Left == nil {
		sum, tilt = findTiltRecursive(root.Right)
		tilt += abs(sum, 0)
		sum += root.Val
		return sum, tilt
	}

	lSum, lTilt := findTiltRecursive(root.Left)
	rSum, rTilt := findTiltRecursive(root.Right)

	tilt = lTilt + rTilt
	tilt += abs(lSum, rSum)
	sum = lSum + rSum + root.Val
	return sum, tilt
}

func abs(a, b int) int {
	val := a - b
	if val < 0 {
		return -val
	}
	return val
}
