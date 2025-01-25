package main

func thirdMax(nums []int) int {
	l := ListManager{}

	for i := range nums {
		l.Add(nums[i])
	}
	return l.getMax()
}

type ListManager struct {
	size int
	root *List
}
type List struct {
	prev *List
	next *List
	Val  int
}

func (l *ListManager) getMax() int {
	if l.size >= 3 {
		return l.root.Val
	}
	cur := l.root

	for {
		if cur.next == nil {
			return cur.Val
		}
		cur = cur.next
	}
}

func (l *ListManager) Add(num int) {
	if l.size == 0 {
		l.root = &List{
			Val: num,
		}
		l.size = 1
		return
	}
	if l.size >= 3 && num < l.root.Val {
		return
	}
	currentNode := l.root
	for {
		if currentNode == nil {
			break
		}
		if num == currentNode.Val {
			break
		}
		if num < currentNode.Val {
			prev := currentNode.prev
			if prev == nil {
				l.root.prev = &List{
					Val:  num,
					next: l.root,
				}
				l.root = l.root.prev
				l.size++
				break
			}
			prev.next = &List{
				Val:  num,
				prev: prev,
				next: currentNode,
			}
			currentNode.prev = prev.next
			l.size++
			break
		}
		if num > currentNode.Val && currentNode.next == nil {
			currentNode.next = &List{
				prev: currentNode,
				next: nil,
				Val:  num,
			}
			l.size++
			break
		}
		currentNode = currentNode.next
	}

	if l.size == 4 {
		l.root = l.root.next
		l.size--
	}
}
