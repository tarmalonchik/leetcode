package main

type LRUCache struct {
	listPos  map[int]*listNode
	capacity int
	size     int
	list     ListManager
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		listPos:  make(map[int]*listNode, capacity),
		list:     ListManager{},
	}
}

func (this *LRUCache) Get(key int) int {
	node, ok := this.listPos[key]
	if !ok {
		return -1
	}
	value := node.Value
	this.list.remove(node)
	this.listPos[key] = this.list.addBack(key, value)
	return value
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.listPos[key]
	if !ok {
		if this.size < this.capacity {
			this.size++
			this.listPos[key] = this.list.addBack(key, value)
			return
		}
		delete(this.listPos, this.list.removeFront().Key)
		this.listPos[key] = this.list.addBack(key, value)
		return
	}
	this.list.remove(node)
	this.listPos[key] = this.list.addBack(key, value)
}

type ListManager struct {
	head *listNode
	tail *listNode
}

type listNode struct {
	Key   int
	Value int
	Prev  *listNode
	Next  *listNode
}

func (l *ListManager) addBack(key, value int) *listNode {
	if l.head == nil && l.tail == nil {
		l.head = &listNode{
			Value: value,
			Key:   key,
		}
		l.tail = l.head
		return l.tail
	}
	l.tail.Next = &listNode{
		Value: value,
		Key:   key,
		Prev:  l.tail,
	}
	l.tail = l.tail.Next
	return l.tail
}

func (l *ListManager) remove(list *listNode) {
	if list == l.head {
		l.removeFront()
		return
	}
	if list == l.tail {
		l.tail = l.tail.Prev
		l.tail.Next = nil
		return
	}
	prev := list.Prev
	next := list.Next
	prev.Next = next
	next.Prev = prev
}

func (l *ListManager) removeFront() *listNode {
	if l.head == nil {
		panic("cannot remove")
	}
	if l.head == l.tail {
		resp := l.head
		l.head = nil
		l.tail = nil
		return resp
	}
	oldHead := l.head
	l.head = l.head.Next
	l.head.Prev = nil
	return oldHead
}
