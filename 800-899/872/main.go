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

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	cont1 := newContainer(root1)
	cont2 := newContainer(root2)

	for {
		n1 := cont1.next()
		n2 := cont2.next()
		if n1 == nil && n2 == nil {
			return true
		}
		if n1 == nil || n2 == nil {
			return false
		}
		if n1.Val == n2.Val {
			continue
		} else {
			return false
		}
	}
}

type container struct {
	currentNode *TreeNode
	direction   int
	nodes       []node
	size        int
}

func newContainer(root *TreeNode) *container {
	return &container{
		currentNode: root,
		direction:   d,
	}
}

func (c *container) next() *TreeNode {
	for {
		if c.currentNode == nil {
			return nil
		}

		if c.currentNode.Right == nil && c.currentNode.Left == nil {
			out := c.currentNode
			c.currentNode, c.direction = c.get()
			return out
		}

		switch c.direction {
		case d:
			if c.currentNode.Left != nil {
				c.add(c.currentNode, l)
				c.currentNode = c.currentNode.Left
				continue
			}
			if c.currentNode.Right != nil {
				c.currentNode = c.currentNode.Right
				continue
			}
			c.currentNode, c.direction = c.get()
			continue
		case l:
			if c.currentNode.Right != nil {
				c.currentNode = c.currentNode.Right
				c.direction = d
			} else {
				c.currentNode, c.direction = c.get()
			}
			continue
		}
	}
}

type node struct {
	val       *TreeNode
	direction int
}

func (c *container) add(in *TreeNode, direction int) {
	if len(c.nodes) > c.size {
		c.nodes[c.size] = node{
			val:       in,
			direction: direction,
		}
		c.size++
		return
	}
	c.nodes = append(c.nodes, node{
		val:       in,
		direction: direction,
	})
	c.size++
}

func (c *container) get() (out *TreeNode, direction int) {
	if c.size == 0 {
		return nil, 0
	}
	item := c.nodes[c.size-1]
	out, direction = item.val, item.direction
	c.size--
	return out, direction
}
