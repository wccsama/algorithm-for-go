package tree

// 递归
func post(root *Tree) {
	if root == nil {
		return
	}
	post(root.Left)
	post(root.Right)
	// print
}

// 非递归
func postNormal(root *Tree) {
	s := Stack{}
	p := root
	for p != nil || !s.empty() {
		for p != nil {
			temp := helpTree{
				root:    p,
				isFirst: true,
			}
			s.push(temp)
			p = p.Left
		}
		if !s.empty() {
			top := s.pop()
			if top.isFirst {
				top.isFirst = false
				s.push(top)
				p = top.root.right
			} else {
				// print
			}
		}
	}
}
