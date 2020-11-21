package tree

// 前序递归
func former(root *Tree) {
	if root == nil {
		return
	}
	// print
	former(root.Left)
	former(root.Right)
}

// 前序非递归
func formerNormal(root *Tree) {
	s := Stack{}
	p := root
	for p != nil || !s.empty() {
		for p != nil {
			// print
			s.push(p)
			p = p.Left
		}
		if !s.empty() {
			p = s.pop()
			p = p.Right
		}
	}
}
