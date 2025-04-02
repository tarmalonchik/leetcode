package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	m := minHolder{}
	recursion(root, &m)
	return m.get()
}

func recursion(root *TreeNode, m *minHolder) {
	if root == nil {
		return
	}
	m.add(root.Val)
	recursion(root.Left, m)
	recursion(root.Right, m)
	return
}

type minHolder struct {
	one      int
	oneValid bool
	two      int
	twoValid bool
}

func (m *minHolder) get() int {
	if !m.oneValid || !m.twoValid {
		return -1
	}
	if m.one == m.two {
		return -1
	}
	if m.one < m.two {
		return m.two
	}
	return m.one
}

func (m *minHolder) add(val int) {
	if !m.oneValid {
		m.one = val
		m.oneValid = true
		return
	}
	if !m.twoValid {
		m.two = val
		m.twoValid = true
		return
	}

	if val == m.one || val == m.two {
		return
	}

	if val < m.one || val < m.two {
		if val < m.one && val < m.two {
			if m.one < m.two {
				m.two = val
			} else {
				m.one = val
			}
			return
		}
		if val < m.one {
			m.one = val
			return
		} else {
			m.two = val
			return
		}
	}

	if m.one == m.two {
		m.one = val
	}
}
