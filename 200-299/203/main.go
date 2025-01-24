package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	var prev *ListNode
	var newHead = head

	for {
		if head == nil {
			break
		}
		if head.Val != val {
			prev = head
			head = head.Next
			continue
		}
		if prev == nil {
			head = head.Next
			newHead = head
			continue
		}
		prev.Next = head.Next
		head = head.Next
		continue
	}
	return newHead
}
