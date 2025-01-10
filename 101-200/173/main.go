package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	nodeStack *nodeStack
}

func Constructor(root *TreeNode) BSTIterator {
	b := BSTIterator{}
	if root == nil {
		return b
	}

	b.addToStack(root, true)
	for {
		if root.Left != nil {
			b.addToStack(root.Left, true)
			root = root.Left
			continue
		}
		break
	}
	return b
}

func (this *BSTIterator) Next() int {
	if this.nodeStack == nil {
		return 0
	}

	resp := this.nodeStack.node.Val

	if this.nodeStack.node.Right != nil {
		rightNode := this.nodeStack.node.Right
		this.getFromStack()
		this.addToStack(rightNode, false)
		for {
			if this.nodeStack.node.Left != nil {
				this.addToStack(this.nodeStack.node.Left, true)
			} else {
				break
			}
		}
		return resp
	}
	this.getFromStack()
	return resp
}

func (this *BSTIterator) HasNext() bool {
	if this.nodeStack == nil {
		return false
	}
	return true
}

func (b *BSTIterator) addToStack(node *TreeNode, isLeft bool) {
	if b.nodeStack == nil {
		b.nodeStack = &nodeStack{
			node:   node,
			prev:   nil,
			isLeft: isLeft,
		}
		return
	}
	b.nodeStack = &nodeStack{
		node:   node,
		prev:   b.nodeStack,
		isLeft: isLeft,
	}
}

func (b *BSTIterator) getFromStack() *TreeNode {
	if b.nodeStack == nil {
		return nil
	}
	resp := b.nodeStack.node
	b.nodeStack = b.nodeStack.prev
	return resp
}

type nodeStack struct {
	prev   *nodeStack
	isLeft bool
	node   *TreeNode
}
