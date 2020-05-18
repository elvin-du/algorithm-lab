//序列化和反序列化二叉树
package code_interviews

import (
	"algorithm-lab/common"
	"fmt"
)

var s = []string{}

func SerializeBT(root *common.BinaryTreeNode) {
	if common.IsNil(root) {
		s = append(s, "$")
		return
	}
	str := fmt.Sprintf("%v", root.Value)
	s = append(s, str)
	SerializeBT(root.Left)
	SerializeBT(root.Right)
}

var index = 0

func Deserialize(str []string) *common.BinaryTreeNode {
	if 0 == len(str) {
		return nil
	}

	if index >= len(str) {
		return nil
	}
	v := str[index]

	if "$" == v {
		index++
		return nil
	}

	root := common.NewBinaryTreeNode(v)
	index++
	root.Left = Deserialize(str)
	root.Right = Deserialize(str)

	return root
}
