package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	mp := make(map[int]interface{})
	return findTargetRecursion(root, k, mp)
}

func findTargetRecursion(root *TreeNode, k int, mp map[int]interface{}) bool {
	if root == nil {
		return false
	}
	_, ok := mp[root.Val]
	if ok {
		return true
	}
	mp[k-root.Val] = nil
	return findTargetRecursion(root.Left, k, mp) || findTargetRecursion(root.Right, k, mp)
}
