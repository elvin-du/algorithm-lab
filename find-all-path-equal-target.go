//剑指offer 二叉树中和为某一值的全部路径
package algorithm_lab

import "algorithm-lab/common"

func FindAllPath(bt *common.BinaryTree, target int) [][]int {
	allPath := [][]int{}
	path := []int{}
	findAllPath(bt.Root, target, 0, &allPath, path)
	return allPath
}

func findAllPath(node *common.BinaryTreeNode, target, curSum int, allPath *[][]int, path []int) {
	path = append(path, node.Value.(int))
	curSum += node.Value.(int)
	isLeaf := common.IsNil(node.Left) && common.IsNil(node.Right)

	//如果为叶子结点， 判断路径的和是否满足要求，满足，保存该路径
	if curSum == target && isLeaf {
		*allPath = append(*allPath, path)
	}

	if !common.IsNil(node.Left) {
		findAllPath(node.Left, target, curSum, allPath, path)
	}

	if !common.IsNil(node.Right) {
		findAllPath(node.Right, target, curSum, allPath, path)
	}

	//如果是叶子结点又不满足要求，退回到父结点，删除当前节点
	path = path[:len(path)-1]
}
