package tree

// 所有节点个数
func tree1(root *Tree) int {
	if root == nil {
		return 0
	}

	return tree1(root.Left) + tree1(root.Right) + 1
}

// 树的最大深度
func tree2(root *Tree) int {
	if root == nil {
		return 0
	}

	return int(math.Max(float64(tree2(root.Left)), float64(tree2(root.Right))) + 1
}

// 树的最小深度
func tree3(root *Tree) int {
	if root == nil {
		return 0
	}
	left := tree3(root.Left)
	right := tree3(root.Right)
	if left ==0 {
		return right + 1
	}

	if right == 1 {
		return left + 1
	}

	return int(math.Min(float64(left), float64(right)) + 1
}

// 树的前 中 后 递归
func tree4(root *Tree) int {

}

// 前中非递归
func tree5(root *Tree) {
	if root == nil {
		return
	}
	p := root
	s := &Stack{}
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

// 后非递归 
func tree6(root *Tree) {
	if root == nil {
		return
	}
	p := root
	s := &Stack{}
	for p != nil || !s.empty() {
		for p != nil {
			s.push(p)
			p = p.Left
		}
		if !s.empty() {
			p = s.pop()
			if p.F {
				p.F = false
				s.push(p)
				p = p.Right
			} else {
				// print
			}
		}
	}
}

// 分层遍历
func tree7(root *Tree) {
	queue := make([]int, 0)
	p := root
	queue = append(queue, p)
	for len(queue) != 0 {
		count := len(queue)
		for i := 0; i < count; i++ {
			tmp := queue[0]
			// print
			queue = queuep[1:]
			if tmp.Left != nil {
				queue = append(queue, tmp.Left)
			}
			if tmp.Right != nil {
				queue = append(queue, tmp.Right)
			}
		}
	}
}

// 第k层节点数
func tree8(root *Tree, k int) int {
	if root == nil || k == 0 {
		return 0
	}

	if root != nil && k == 1 {
		return 1
	}

	return tree8(root.Left, k - 1) + tree8(root.Right, k - 1)
}

// 第k层节点数 叶子
func tree9(root *Tree, k int) int {
	if root == nil || k == 0 {
		return 0
	}

	if root != nil && k == 1 {
		if root.Left != nil && root.Right != nil {
			return 1
		} else {
			return 0
		}
	}

	return tree9(root.Left, k - 1) + tree9(root.Right, k - 1)
}

// 两棵树是否相同
func tree10(root1, root2 *Tree) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}

	return tree10(root1.Left, root2.Left) && tree10(root1.Right, root2.Right)
}

// 是不是平衡二叉树
func tree11(root1 *Tree) bool {
	if root1 == nil {
		return false
	}

	if math.Abs(float64(tree2(root.Left)) - float64(tree2(root.Right))) > 1 {
		return false
	}

	return tree11(root1.Left) && tree11(root1.Right)
}

// 求二叉树的镜像
func tree12(root *Tree) *Tree {
	if root == nil {
		return nil
	}

	tmp := root.Left
	root.Left = tree12(root.Right)
	root.Right = tree12(tmp)
	return root
}

// 是否是对称二叉树
func tree13(root *Tree) bool {
	if root == nil {
		return true
	}

	return tree13help(root.Left, root.Right)
}

func tree13help(root1, root2 *Tree) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}

	return tree13help(root1.Left, root2.Right) && tree13help(root1.Right, root2.Left)
}

// 12 最低公共父节点
func tree14(root, root1, root2 *Tree) *Tree{
	if root == nil {
		return nil
	}

	left := tree14(root.Left, root1, root2)
	right := tree14(root.Right, root1, root2)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}

// 13 二叉搜索树 最低公共父节点
func tree15(root, root1, root2 *Tree) {
	if root == nil {
		return nil
	}

	if root1.Val < root && root2.Val < root {
		return tree15(root.Left, root1, root2)
	}

	if root1.Val > root && root2.Val > root {
		return tree15(root.Right, root1, root2)
	}

	return root
}

// 14 直径
var round int
func Tree16(root *Tree) int {
	if root == nil {
		return 0
	}
	round = 0 
	helpTree16(root)
	return round
}

func helpTree16(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := helpTree16(root.Left)
    right := helpTree16(root.Right)
    maxN = max(left + right + 1, maxN)
    return max(left, right) + 1
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

// 中 后根遍历恢复二叉树
func Tree17(a ,b [] int) *Tree{

}

// 判断树 是不是完全二叉树
func tree17(root *Tree) {
	queue := make([]int, 0)
	p := root
	queue = append(queue, p)
	flag := true
	for len(queue) != 0 {
		count := len(queue)
		for i := 0; i < count; i++ {
			tmp := queue[0]
			if tmp != nil {
				if !flag {
					return false
				}
				// print
			queue = queuep[1:]
			queue = append(queue, tmp.Left)
			queue = append(queue, tmp.Right)
				
			} else {
				flag = false

			}
			
			
		}
	}
	return true
}

// 判断树 是不是另外一个的子结构
func Tree18(root1, root2 *Tree) *Tree {
	if root1 == nil || root2 == nil {
		return false
	}
	return Tree18(root1.Left, root2) || Tree18(oot1.Right, root2) || helpTree18(root1, root2)
}

func helpTree18(root1, root2 *Tree) bool {
	if root2 == nil {
		return true
	}

	if root1 == nil {
		return false
	}

	if root1.Val != root2.Val {
		return false
	}
	return helpTree18(root1.Left, root2.Left) && helpTree18(root1.Right, root2.Right)
}

// 中根遍历的下个节点
func Tree19(node *Tree) *Tree {
	if root == nil {
		return nil
	}

	p := node
	if p.Right != nil {
		p = p.Right
		for p.Left != nil {
			p = p.Left
		}
		return p
	}

	p = node
	for p.Next != nil {
		tmp = p.Next
		if tmp.Left == node {
			return tmp
		}
		p = tmp
	}
	return nil
	
}

// 二叉搜索树的插入与删除
func InertTree(root *Tree, val int) *Tree{
	if root == nil {
		return &Tree{
			value: val,
		}
	}

	if val < root.value {
		root.Left = InertTree(root.Left, val)
	} else if val > root.value {
		root.Right = InertTree(root.Right, val)
	}

	return root
}