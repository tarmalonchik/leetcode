package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	newRoot := &TreeNode{}
	if root1 == nil && root2 == nil {
		return nil
	} else if root1 != nil && root2 != nil {
		newRoot.Val = root1.Val + root2.Val
		newRoot.Left = mergeTrees(root1.Left, root2.Left)
		newRoot.Right = mergeTrees(root1.Right, root2.Right)
	} else {
		if root1 != nil {
			newRoot.Val = root1.Val
			newRoot.Left = mergeTrees(root1.Left, nil)
			newRoot.Right = mergeTrees(root1.Right, nil)
		} else {
			newRoot.Val = root2.Val
			newRoot.Left = mergeTrees(nil, root2.Left)
			newRoot.Right = mergeTrees(nil, root2.Right)
		}
	}
	return newRoot
}
