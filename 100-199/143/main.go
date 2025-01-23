package main

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	head1, head2 := divideTwoParts(head)
	head2 = reverseList(head2)
	for {
		if head1 == nil || head2 == nil {
			break
		}
		nxt1 := head1.Next
		nxt2 := head2.Next
		head1.Next = head2
		head1.Next.Next = nxt1
		head1 = nxt1
		head2 = nxt2
	}
}

func divideTwoParts(head *ListNode) (*ListNode, *ListNode) {
	lazy := head
	fast := head
	for {
		if fast.Next == nil || fast.Next.Next == nil {
			second := lazy.Next
			lazy.Next = nil
			return head, second
		}
		lazy = lazy.Next
		fast = fast.Next.Next
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

type ListNode struct {
	Val  int
	Next *ListNode
}
