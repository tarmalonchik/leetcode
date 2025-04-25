package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{
			Val: val,
		}
		return root
	}
	current := root
	for {
		if val > current.Val {
			if current.Right == nil {
				current.Right = &TreeNode{
					Val: val,
				}
				return root
			}
			current = current.Right
		} else {
			if current.Left == nil {
				current.Left = &TreeNode{
					Val: val,
				}
				return root
			}
			current = current.Left
		}
	}
}
