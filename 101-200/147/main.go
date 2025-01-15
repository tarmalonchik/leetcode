package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func insert(root *ListNode, val int, stop *ListNode) *ListNode {
	current := root

	if val < current.Val {
		return &ListNode{
			Val:  val,
			Next: current,
		}
	}

	for {
		if current.Next == nil || current.Next == stop {
			if current.Next == nil {
				current.Next = &ListNode{
					Val: val,
				}
				return root
			}
			current.Next = &ListNode{
				Val:  val,
				Next: stop,
			}
			return root
		}
		if val <= current.Next.Val {
			item := current.Next
			current.Next = &ListNode{
				Val:  val,
				Next: item,
			}
			return root
		}
		current = current.Next
	}
}

func insertionSortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	newHead := head
	current := head

	for {
		if current == nil || current.Next == nil {
			break
		}
		if current.Next.Val < current.Val {
			nxt := current.Next
			current.Next = nxt.Next
			newHead = insert(newHead, nxt.Val, current.Next)
			continue
		}
		current = current.Next
	}
	return newHead
}
