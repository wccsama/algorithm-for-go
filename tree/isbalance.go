package tree

import "math"

// 根据最大深度算
func isBalance(root *Tree) bool {
	if root == nil {
		return true
	}
	return math.Abs(treeMax(right)-treeMax(left)) > 1 && isBalance(left) && isBalance(right)
}
