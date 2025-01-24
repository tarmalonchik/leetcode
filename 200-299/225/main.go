package main

type MyStack struct {
	lastNode *ListItem
}

func Constructor() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	if this.lastNode == nil {
		this.lastNode = &ListItem{
			Val: x,
		}
		return
	}
	item := &ListItem{
		Val:  x,
		Prev: this.lastNode,
	}
	this.lastNode = item
}

func (this *MyStack) Pop() int {
	if this.lastNode == nil {
		return 0
	}
	out := this.lastNode
	this.lastNode = this.lastNode.Prev
	return out.Val
}

func (this *MyStack) Top() int {
	if this.lastNode == nil {
		return 0
	}
	return this.lastNode.Val
}

func (this *MyStack) Empty() bool {
	return this.lastNode == nil
}

type ListItem struct {
	Val  int
	Prev *ListItem
}
