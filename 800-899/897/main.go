package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const (
	d = iota
	l = iota
)

func increasingBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	s := &stack{}
	var newRoot *TreeNode
	var tail *TreeNode
	direction := d

	for {
		if root == nil {
			break
		}
		switch direction {
		case d:
			if root.Left != nil {
				s.add(root, l)
				root = root.Left
				continue
			}
			if root.Right != nil {
				if tail == nil {
					tail = addToTail(nil, root.Val)
					newRoot = tail
				} else {
					tail = addToTail(tail, root.Val)
				}
				root = root.Right
				continue
			}
			if tail == nil {
				tail = addToTail(nil, root.Val)
				newRoot = tail
			} else {
				tail = addToTail(tail, root.Val)
			}
			root, direction = s.get()
			continue
		case l:
			tail = addToTail(tail, root.Val)
			if root.Right != nil {
				root = root.Right
				direction = d
			} else {
				root, direction = s.get()
			}
			continue
		}
	}
	return newRoot
}

func addToTail(tail *TreeNode, val int) *TreeNode {
	if tail == nil {
		return &TreeNode{
			Val: val,
		}
	}
	tail.Right = &TreeNode{
		Val: val,
	}
	return tail.Right
}

type stack struct {
	next *node
}

type node struct {
	next      *node
	item      *TreeNode
	direction int
}

func (s *stack) add(treeNode *TreeNode, direction int) {
	if s.next == nil {
		s.next = &node{
			item:      treeNode,
			direction: direction,
		}
		return
	}
	newItem := &node{
		item:      treeNode,
		next:      s.next,
		direction: direction,
	}
	s.next = newItem
}

func (s *stack) get() (*TreeNode, int) {
	val := s.next
	if s.next != nil {
		s.next = s.next.next
	}
	if val == nil {
		return nil, 0
	}
	return val.item, val.direction
}
