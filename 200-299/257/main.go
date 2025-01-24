package main

import (
	"fmt"
	"strconv"
)

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	fmt.Println(binaryTreePaths(root))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePaths(root *TreeNode) []string {
	return recursion("", root)
}

func recursion(prefix string, node *TreeNode) []string {
	if prefix == "" {
		prefix += strconv.Itoa(node.Val)
	} else {
		prefix += "->" + strconv.Itoa(node.Val)
	}

	out := make([]string, 0)

	if node.Left != nil {
		out = append(out, recursion(prefix, node.Left)...)
	}
	if node.Right != nil {
		out = append(out, recursion(prefix, node.Right)...)
	}
	if node.Left == nil && node.Right == nil {
		return []string{prefix}
	}
	return out
}
