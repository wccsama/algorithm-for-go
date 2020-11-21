package tree

var result []int
var temp []int

func sum(root *Tree, sum int) []int {
	if root == nil {
		return result
	}

	temp = append(temp, root.value)
	sum -= root.value

	if sum == 0 && root.Left == nil && root.Right == nil {
		result = append(result, temp...)
	} else {
		FindPath(root.left, sum)
		FindPath(root.right, sum)
	}

	temp = temp[:len(temp)]
	return result

}

var result = make([][]int, 0)

func pathSum(root *TreeNode, sum int) [][]int {
	FindPath(root, sum)
	return result
}

func FindPath(root *TreeNode, sum int) {
	if root == nil {
		return
	}

	sum -= root.Val
	tmp = append(tmp, root.Val)
	if sum == 0 && root.Left == nil && root.Right == nil {
		cop := make([]int, tmp)
		cop = copy(tmp)
		result = append(result, cop)
	} else {
		FindPath(root.left, sum)
		FindPath(root.right, sum)
	}

	sum += tem[len(tem)-1]
	tem = tmp[:len(tem)]
}


var res [][]int
func pathSum(root *TreeNode, sum int) [][]int {
    res = make([][]int,0)
    if root == nil {
        return res
    }
   
    helper(root,sum,[]int{})
    return res
}

func helper(root *TreeNode,sum int,path []int) {
    if root == nil {
        return
    }
    path = append(path,root.Val)
    sum -= root.Val
    if root.Left == nil && root.Right == nil {
        if sum == 0 {
            tmp := make([]int,len(path))
            copy(tmp,path)
            res = append(res,tmp)
        }
    }

    helper(root.Left,sum,path)
    helper(root.Right,sum,path)
}


