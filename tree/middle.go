package tree

// 中序递归
func middle(root *Tree) {
	if root == nil {
		return
	}
	middle(root.Left)
	// print
	middle(root.Right)
}

// 中序非递归
func middleNormal(root *Tree) {
	s := Stack{}
	p := root
	for p != nil || !s.empty() {
		for p != nil {
			s.push(p)
			p = p.Left
		}
		if !s.empty() {
			p = s.pop()
			// print
			p = p.Right
		}
	}
}

func model(root *Tree) {
	stack := Stack{}
	p := root
	for p != nil || !s.empty() {
		for p != nil {
			s.push(p)
			// print
			p = p.Left
		}

		for !s.empty() {
			p = s.pop()
			// print
			p = p.Right
		}
	}
}
