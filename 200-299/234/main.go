package main

func isPalindrome(head *ListNode) bool {
	if head.Next == nil {
		return true
	}
	a, b := divideTwo(head)
	a = reverseList(a)

	for {
		if a == nil {
			return true
		}
		if a.Val != b.Val {
			return false
		}
		a = a.Next
		b = b.Next
	}
}

func reverseList(head *ListNode) *ListNode {
	var tail *ListNode
	for {
		if head.Next == nil {
			head.Next = tail
			return head
		}
		next := head.Next
		head.Next = tail
		tail = head
		head = next
	}
}

func divideTwo(a *ListNode) (*ListNode, *ListNode) {
	if a == nil {
		return nil, nil
	}
	var prevLazy *ListNode
	lazy := a
	fast := a
	for {
		if fast.Next == nil || fast.Next.Next == nil {
			b := a
			c := lazy.Next
			lazy.Next = nil
			if fast.Next == nil && prevLazy != nil {
				prevLazy.Next = nil
			}
			return b, c
		}
		prevLazy = lazy
		lazy = lazy.Next
		fast = fast.Next.Next
	}
}

type ListNode struct {
	Val  int
	Next *ListNode
}
