package tree

func next(des *Tree) *Tree {
	if des == nil {
		return nil
	}

	if des.Right != nil {
		root := des.Right
		for root.left != nil {
			root := root.left
		}
		return root
	}

	for des.F != nil {
		if des.F.Left == des {
			return des.F
		} else {
			des = des.F
		}

	}

}
