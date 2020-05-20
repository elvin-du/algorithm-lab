package code_interviews

import "algorithm-lab/common"

var (
	target  = 10
	path    = []int{}
	allpath = [][]int{}
)

func FindPath(root *common.BinaryTreeNode) {
	if common.IsNil(root) {
		return
	}

	path = append(path, root.Value.(int))
	target = target - root.Value.(int)
	if 0 == target && common.IsNil(root.Left) && common.IsNil(root.Right) {
		tmp := make([]int, len(path))
		copy(tmp, path)
		allpath = append(allpath, tmp)
	}

	FindPath(root.Left)
	FindPath(root.Right)

	path = path[:len(path)-1]
	target -= root.Value.(int)
}
