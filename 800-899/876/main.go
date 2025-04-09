package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	singleStep := head
	doubleStep := head

	for {
		if doubleStep.Next == nil {
			return singleStep
		}
		if doubleStep.Next.Next == nil {
			return singleStep.Next
		}
		singleStep = singleStep.Next
		doubleStep = doubleStep.Next.Next
	}
}
