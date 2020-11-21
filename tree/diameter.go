package tree

import "math"

var path int

func diameter(root *Tree) int {
	if root == nil {
		return 0
	}
	_ = getD(root)
	return path
}

func getD(root *Tree) int {
	if root == nil {
		return 0
	}
	left := getD(root.Left)
	right := getD(root.Right)
	path := math.Max(path, left+right)
	return math.Max(left, right) + 1

}
