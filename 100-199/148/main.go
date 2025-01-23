package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if head.Next.Next == nil {
		return sortListWithLessThanTwo(head)
	}

	a, b := divideTwo(head)
	a = sortList(a)
	b = sortList(b)
	return mergeSortedLists(a, b)
}

func mergeSortedLists(a, b *ListNode) *ListNode {
	if a == nil {
		return b
	} else if b == nil {
		return a
	}

	var mainNode *ListNode
	var secondNode *ListNode

	if a.Val <= b.Val {
		mainNode = a
		secondNode = b
	} else {
		mainNode = b
		secondNode = a
	}
	head := mainNode

	for {
		if mainNode.Next == nil {
			mainNode.Next = secondNode
			break
		}
		if secondNode == nil {
			break
		}
		if mainNode.Next.Val <= secondNode.Val {
			mainNode = mainNode.Next
			continue
		} else {
			insertNode := secondNode
			secondNode = secondNode.Next
			mainNext := mainNode.Next
			mainNode.Next = insertNode
			insertNode.Next = mainNext
			mainNode = insertNode
		}
	}
	return head
}

func sortListWithLessThanTwo(a *ListNode) *ListNode {
	if a.Val > a.Next.Val {
		b := a
		a = a.Next
		b.Next = nil
		a.Next = b
	}
	return a
}

func divideTwo(a *ListNode) (*ListNode, *ListNode) {
	if a == nil {
		return nil, nil
	}
	lazy := a
	fast := a
	for {
		if fast.Next == nil || fast.Next.Next == nil {
			b := a
			c := lazy.Next
			lazy.Next = nil
			return b, c
		}
		lazy = lazy.Next
		fast = fast.Next.Next
	}
}
