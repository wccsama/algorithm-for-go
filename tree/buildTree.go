package tree

// 根据前序遍历、中序遍历获取数
func buildTree1(pre, mid []int) *Tree {
	if len(pre) == 0 || len(mid) == 0 {
		return nil
	}
	return build(pre, 0, len(pre)-1, mid, 0, len(mid)-1)
}

func build(pre []int, preS, preE int, mid []int, midS, midE int) *Tree {
	if preS < preE || midS < midE {
		return nil
	}

	tree := &Tree{
		value: pre[preS],
	}

	for i := midS; i <= midE; i++ {
		if mid[i] == pre[preS] {
			tree.Left = build(pre, preS+1, preS+i-midS, mid, midS, i-1)
			tree.Right = build(pre, preS+i-midS+1, preE, mid, i+1, midE)
		}
	}
	return tree
}
