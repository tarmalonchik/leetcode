package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findMode(root *TreeNode) []int {
	maxNum := 1
	out := make([]int, 0)

	mp := make(map[int]int)
	findModeRec(root, mp, &maxNum)

	for key, val := range mp {
		if val == maxNum {
			out = append(out, key)
		}
	}
	return out
}

func findModeRec(root *TreeNode, mp map[int]int, maxNum *int) {
	mp[root.Val]++

	if mp[root.Val] > *maxNum {
		*maxNum = mp[root.Val]
	}

	if root.Left != nil {
		findModeRec(root.Left, mp, maxNum)
	}
	if root.Right != nil {
		findModeRec(root.Right, mp, maxNum)
	}
}
